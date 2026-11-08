package main

import (
	"bytes"
	"cmp"
	"encoding/json"
	"flag"
	"go/ast"
	"go/constant"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"unicode"
)

type apiField struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Default string   `json:"default"`
	Doc     string   `json:"doc"`
	Values  []string `json:"values,omitempty"`
}

type apiStruct struct {
	Fields []apiField `json:"fields"`
	Page   string     `json:"-"`
}

func main() {
	out := flag.String("out", "catalog", "catalog directory to write <page>/api.json into")
	flag.Parse()
	dirs := flag.Args()
	if len(dirs) == 0 {
		found, err := filepath.Glob("pkg/*")
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range found {
			if info, err := os.Stat(d); err == nil && info.IsDir() && filepath.Base(d) != "internal" {
				dirs = append(dirs, d)
			}
		}
	}

	api, err := extract(dirs...)
	if err != nil {
		log.Fatal(err)
	}
	merged := pagesOf(*out, api)
	for dir, structs := range merged {
		data, err := json.MarshalIndent(structs, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		if err := writeIfChanged(filepath.Join(dir, "api.json"), append(data, '\n')); err != nil {
			log.Fatal(err)
		}
	}
}

// pagesOf groups the structs by the catalog page directory each belongs to, and adds to a page
// the structs its index.mdx names in a GoApi, as date-picker shows CalendarProps.
func pagesOf(out string, api map[string]apiStruct) map[string]map[string]apiStruct {
	merged := map[string]map[string]apiStruct{}
	add := func(dir, name string, s apiStruct) {
		if merged[dir] == nil {
			merged[dir] = map[string]apiStruct{}
		}
		merged[dir][name] = s
	}
	for name, s := range api {
		if dir, ok := pageDir(out, s.Page); ok {
			add(dir, name, s)
		}
	}
	pages, _ := filepath.Glob(filepath.Join(out, "*", "index.mdx"))
	for _, mdx := range pages {
		data, err := os.ReadFile(mdx)
		if err != nil {
			continue
		}
		for _, m := range goAPIName.FindAllSubmatch(data, -1) {
			if s, ok := api[string(m[1])]; ok {
				add(filepath.Dir(mdx), string(m[1]), s)
			}
		}
	}
	return merged
}

var goAPIName = regexp.MustCompile(`<GoApi[^>]*\bname="(\w+)"`)

// pageDir finds the catalog page a struct's page name belongs to: the page itself, or the nearest
// shorter one, as toast for ToastHostProps.
func pageDir(out, page string) (string, bool) {
	for {
		dir := filepath.Join(out, page)
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return dir, true
		}
		i := strings.LastIndexByte(page, '-')
		if i < 0 {
			return "", false
		}
		page = page[:i]
	}
}

func writeIfChanged(file string, data []byte) error {
	if old, err := os.ReadFile(file); err == nil && bytes.Equal(old, data) {
		return nil
	}
	return os.WriteFile(file, data, 0o644)
}

func pageOf(structName string) string {
	name := []rune(strings.TrimSuffix(structName, "Props"))
	var b strings.Builder
	for i, r := range name {
		upper := unicode.IsUpper(r)
		if upper && i > 0 {
			prevLower := unicode.IsLower(name[i-1])
			nextLower := i+1 < len(name) && unicode.IsLower(name[i+1])
			if prevLower || (unicode.IsUpper(name[i-1]) && nextLower) {
				b.WriteByte('-')
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

type parsedPackage struct {
	name       string
	underlying map[string]string
	props      []*ast.TypeSpec
	structs    map[string]*ast.TypeSpec
	values     map[string][]string
	zero       map[string]string
}

func extract(dirs ...string) (map[string]apiStruct, error) {
	var pkgs []*parsedPackage
	byName := map[string]*parsedPackage{}
	for _, dir := range dirs {
		pkg, err := parseDir(dir)
		if err != nil {
			return nil, err
		}
		if pkg != nil {
			pkgs = append(pkgs, pkg)
			byName[pkg.name] = pkg
		}
	}

	api := map[string]apiStruct{}
	for _, pkg := range pkgs {
		for _, ts := range pkg.props {
			page := pageOf(ts.Name.Name)
			api[ts.Name.Name] = apiStruct{Fields: pkg.fieldsOf(ts, byName), Page: page}
			for _, f := range ts.Type.(*ast.StructType).Fields.List {
				elem := strings.TrimLeft(types.ExprString(f.Type), "[]*")
				if nested, ok := pkg.structs[elem]; ok && !strings.HasSuffix(elem, "Props") {
					api[elem] = apiStruct{Fields: pkg.fieldsOf(nested, byName), Page: page}
				}
			}
		}
	}
	return api, nil
}

func (p *parsedPackage) fieldsOf(ts *ast.TypeSpec, byName map[string]*parsedPackage) []apiField {
	var fields []apiField
	for _, f := range ts.Type.(*ast.StructType).Fields.List {
		typ := types.ExprString(f.Type)
		owner, local := p, typ
		if before, after, ok := strings.Cut(typ, "."); ok && byName[before] != nil {
			owner, local = byName[before], after
		}
		shown := typ
		if _, declared := owner.underlying[local]; declared && owner == p {
			shown = p.name + "." + typ
		}
		for _, name := range f.Names {
			if !name.IsExported() {
				continue
			}
			fields = append(fields, apiField{
				Name:    name.Name,
				Type:    shown,
				Default: owner.defaultOf(local),
				Doc:     strings.TrimSpace(f.Doc.Text()),
				Values:  owner.qualified(local),
			})
		}
	}
	return fields
}

func (p *parsedPackage) qualified(typ string) []string {
	var out []string
	for _, v := range p.values[typ] {
		out = append(out, p.name+"."+v)
	}
	return out
}

func parseDir(dir string) (*parsedPackage, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return nil, err
	}
	slices.Sort(files)

	fset := token.NewFileSet()
	pkg := &parsedPackage{underlying: map[string]string{}, structs: map[string]*ast.TypeSpec{}}
	var parsed []*ast.File
	for _, file := range files {
		if strings.HasSuffix(file, "_test.go") {
			continue
		}
		f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, f)
		pkg.name = f.Name.Name
		for _, decl := range f.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.TYPE {
				continue
			}
			for _, spec := range gen.Specs {
				ts := spec.(*ast.TypeSpec)
				pkg.underlying[ts.Name.Name] = types.ExprString(ts.Type)
				if _, isStruct := ts.Type.(*ast.StructType); isStruct && ts.Name.IsExported() {
					pkg.structs[ts.Name.Name] = ts
					if strings.HasSuffix(ts.Name.Name, "Props") {
						pkg.props = append(pkg.props, ts)
					}
				}
			}
		}
	}
	if len(parsed) == 0 {
		return nil, nil
	}
	pkg.values, pkg.zero = collectEnums(fset, parsed)
	return pkg, nil
}

type stubImporter struct{}

func (stubImporter) Import(importPath string) (*types.Package, error) {
	pkg := types.NewPackage(importPath, path.Base(importPath))
	pkg.MarkComplete()
	return pkg, nil
}

type enumConst struct {
	name  string
	value int64
}

func collectEnums(fset *token.FileSet, files []*ast.File) (values map[string][]string, zero map[string]string) {
	conf := types.Config{Importer: stubImporter{}, Error: func(error) {}}
	pkg, _ := conf.Check(files[0].Name.Name, fset, files, nil)

	byType := map[string][]enumConst{}
	for _, name := range pkg.Scope().Names() {
		c, ok := pkg.Scope().Lookup(name).(*types.Const)
		if !ok || !c.Exported() {
			continue
		}
		named, ok := c.Type().(*types.Named)
		if !ok || named.Obj().Pkg() != pkg {
			continue
		}
		basic, ok := named.Underlying().(*types.Basic)
		if !ok || basic.Info()&types.IsInteger == 0 {
			continue
		}
		v, exact := constant.Int64Val(c.Val())
		if !exact {
			continue
		}
		byType[named.Obj().Name()] = append(byType[named.Obj().Name()], enumConst{name, v})
	}

	values, zero = map[string][]string{}, map[string]string{}
	for typ, consts := range byType {
		slices.SortStableFunc(consts, func(a, b enumConst) int { return cmp.Compare(a.value, b.value) })
		for _, c := range consts {
			values[typ] = append(values[typ], c.name)
			if c.value == 0 {
				zero[typ] = c.name
			}
		}
	}
	return values, zero
}

func (p *parsedPackage) defaultOf(typ string) string {
	if name, ok := p.zero[typ]; ok {
		return p.name + "." + name
	}
	for seen := 0; seen < 8; seen++ {
		next, ok := p.underlying[typ]
		if !ok {
			break
		}
		typ = next
	}
	switch {
	case typ == "bool":
		return "false"
	case typ == "string":
		return `""`
	case strings.HasPrefix(typ, "int"), strings.HasPrefix(typ, "uint"), strings.HasPrefix(typ, "float"), typ == "byte", typ == "rune":
		return "0"
	case strings.HasPrefix(typ, "*"), strings.HasPrefix(typ, "[]"), strings.HasPrefix(typ, "map["), strings.HasPrefix(typ, "func"), strings.HasPrefix(typ, "chan "):
		return "nil"
	}
	return ""
}
