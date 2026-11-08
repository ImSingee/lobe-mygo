package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
	"github.com/lobehub/lobe-mygo/catalog/registry"
)

func TestCatalogRegistersEveryDemo(t *testing.T) {
	want := []string{"button/basic", "button/sizes", "button/variants", "tokens/swatches"}
	for _, name := range want {
		if _, ok := registry.Get(name); !ok {
			t.Errorf("demo %s not registered; have %v", name, registry.Names())
		}
	}
}

func TestEveryDemoRendersInThePreview(t *testing.T) {
	for _, name := range registry.Names() {
		view, _ := registry.Get(name)
		tt := ui.NewTester(func(c *ui.Context) { preview(c, view) }, 720, 400)
		if strings.HasPrefix(name, "button/") {
			if _, ok := tt.AccessByRole(ui.RoleButton, ""); !ok {
				t.Errorf("%s shows no button; texts %q", name, tt.Texts())
			}
		}
	}
}

func TestPreviewPaintsContainerSurface(t *testing.T) {
	view, _ := registry.Get("button/basic")
	tt := ui.NewTester(func(c *ui.Context) { preview(c, view) }, 320, 200)
	img := tt.Image()
	scale := img.Bounds().Dx() / 320
	px := img.RGBAAt(4*scale, 4*scale)
	if got, want := ui.RGB(px.R, px.G, px.B), lobe.Light().BgContainer; got != want {
		t.Errorf("corner = %+v, want BgContainer %+v", got, want)
	}
}

func TestSwatchesListSemanticTokens(t *testing.T) {
	view, _ := registry.Get("tokens/swatches")
	tt := ui.NewTester(func(c *ui.Context) { preview(c, view) }, 720, 2000)
	for _, s := range []string{"Text", "#1f1f1f", "Primary"} {
		if !tt.HasText(s) {
			t.Errorf("swatches lack %q", s)
		}
	}
}

func TestSwatchesScrollInsideAShortFrame(t *testing.T) {
	view, _ := registry.Get("tokens/swatches")
	tt := ui.NewTester(func(c *ui.Context) { preview(c, view) }, 884, 640)
	before, ok := tt.Find("Fill")
	if !ok {
		t.Fatalf("no Fill heading; texts %q", tt.Texts())
	}
	tt.Scroll(442, 320, 0, 400)
	tt.Advance(time.Second)
	after, _ := tt.Find("Fill")
	if after.Y >= before.Y {
		t.Errorf("Fill heading at y=%v before and y=%v after scrolling down", before.Y, after.Y)
	}
}

func TestEveryGoDemoNameInMDXIsRegistered(t *testing.T) {
	pages, err := filepath.Glob("*/index.mdx")
	if err != nil || len(pages) == 0 {
		t.Fatalf("no catalog pages found: %v", err)
	}
	pattern := regexp.MustCompile(`<GoDemo[^>]*\bname="([^"]+)"`)
	for _, page := range pages {
		src, err := os.ReadFile(page)
		if err != nil {
			t.Fatal(err)
		}
		matches := pattern.FindAllStringSubmatch(string(src), -1)
		if len(matches) == 0 {
			t.Errorf("%s has no <GoDemo name=...>", page)
		}
		for _, m := range matches {
			if _, ok := registry.Get(m[1]); !ok {
				t.Errorf("%s uses demo %q, which is not registered", page, m[1])
			}
		}
	}
}
