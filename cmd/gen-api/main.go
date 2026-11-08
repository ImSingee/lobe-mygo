package main

import (
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
	pages := map[string]map[string]apiStruct{}
	for name, s := range api {
		page := pageOf(name)
		if pages[page] == nil {
			pages[page] = map[string]apiStruct{}
		}
		pages[page][name] = s
	}
	for page, structs := range pages {
		pageDir := filepath.Join(*out, page)
		if _, err := os.Stat(pageDir); err != nil {
			continue
		}
		data, err := json.MarshalIndent(structs, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(pageDir, "api.json"), append(data, '\n'), 0o644); err != nil {
			log.Fatal(err)
		}
	}
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
			var s apiStruct
			for _, f := range ts.Type.(*ast.StructType).Fields.List {
				typ := types.ExprString(f.Type)
				owner, local := pkg, typ
				if before, after, ok := strings.Cut(typ, "."); ok && byName[before] != nil {
					owner, local = byName[before], after
				}
				shown := typ
				if _, declared := owner.underlying[local]; declared && owner == pkg {
					shown = pkg.name + "." + typ
				}
				for _, name := range f.Names {
					if !name.IsExported() {
						continue
					}
					s.Fields = append(s.Fields, apiField{
						Name:    name.Name,
						Type:    shown,
						Default: owner.defaultOf(local),
						Doc:     strings.TrimSpace(f.Doc.Text()),
						Values:  owner.qualified(local),
					})
				}
			}
			api[ts.Name.Name] = s
		}
	}
	return api, nil
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
	pkg := &parsedPackage{underlying: map[string]string{}}
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
				if _, isStruct := ts.Type.(*ast.StructType); isStruct && ts.Name.IsExported() && strings.HasSuffix(ts.Name.Name, "Props") {
					pkg.props = append(pkg.props, ts)
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
