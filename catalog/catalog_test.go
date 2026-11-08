package main

import (
	"github.com/lobehub/lobe-mygo/pkg/theme"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
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
	if got, want := ui.RGB(px.R, px.G, px.B), theme.Light().BgContainer; got != want {
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

func TestOverlayDemosOpen(t *testing.T) {
	for _, c := range []struct{ demo, trigger, shows string }{
		{"modal/basic", "Open modal", "Publish changes"},
		{"popover/basic", "Share", "Copy link"},
		{"dropdown-menu/basic", "Actions", "Duplicate"},
	} {
		view, ok := registry.Get(c.demo)
		if !ok {
			t.Errorf("%s not registered", c.demo)
			continue
		}
		tt := ui.NewTester(func(ctx *ui.Context) { preview(ctx, view) }, 720, 400)
		r, ok := tt.Find(c.trigger)
		if !ok {
			t.Errorf("%s: no %q; texts %q", c.demo, c.trigger, tt.Texts())
			continue
		}
		tt.Press(r.X+r.W/2, r.Y+r.H/2)
		tt.Release(r.X+r.W/2, r.Y+r.H/2)
		tt.Advance(time.Second)
		if !tt.HasText(c.shows) {
			t.Errorf("%s: %q not shown after pressing %q; texts %q", c.demo, c.shows, c.trigger, tt.Texts())
		}
	}
	for _, demo := range []string{"tooltip/basic", "tooltip/placement", "popover/placement", "dropdown-menu/action-icon", "modal/texts"} {
		if _, ok := registry.Get(demo); !ok {
			t.Errorf("%s not registered", demo)
		}
	}
}

func TestLayoutDemosRegistered(t *testing.T) {
	for _, demo := range []string{"divider/basic", "divider/vertical", "grid/basic", "scroll-area/basic", "scroll-area/horizontal", "draggable-panel/basic", "draggable-panel/placement"} {
		if _, ok := registry.Get(demo); !ok {
			t.Errorf("%s not registered", demo)
		}
	}
}

func TestFeedbackDemos(t *testing.T) {
	for _, demo := range []string{"alert/basic", "alert/variants", "alert/description", "drawer/basic", "drawer/placement", "toast/basic", "toast/action"} {
		if _, ok := registry.Get(demo); !ok {
			t.Errorf("%s not registered", demo)
		}
	}
	for _, c := range []struct{ demo, trigger, shows string }{
		{"drawer/basic", "Open drawer", "Notification settings"},
		{"toast/basic", "Success", "Changes saved"},
	} {
		view, ok := registry.Get(c.demo)
		if !ok {
			continue
		}
		tt := ui.NewTester(func(ctx *ui.Context) { preview(ctx, view) }, 720, 400)
		r, ok := tt.Find(c.trigger)
		if !ok {
			t.Errorf("%s: no %q; texts %q", c.demo, c.trigger, tt.Texts())
			continue
		}
		tt.ClickAt(r.X+r.W/2, r.Y+r.H/2)
		tt.Advance(time.Second)
		if !tt.HasText(c.shows) {
			t.Errorf("%s: %q not shown after clicking %q", c.demo, c.shows, c.trigger)
		}
	}
}

func TestDataDisplayDemos(t *testing.T) {
	for _, demo := range []string{
		"hotkey/basic", "hotkey/variants", "avatar/basic", "avatar/fallback", "avatar/sizes",
		"skeleton/basic", "skeleton/animation", "empty/basic", "empty/page",
		"accordion/basic", "accordion/variants",
	} {
		if _, ok := registry.Get(demo); !ok {
			t.Errorf("%s not registered", demo)
		}
	}
	view, ok := registry.Get("accordion/basic")
	if !ok {
		return
	}
	tt := ui.NewTester(func(ctx *ui.Context) { preview(ctx, view) }, 720, 400)
	if err := tt.Click("What is Lobe UI?"); err != nil {
		t.Fatalf("accordion/basic: %v; texts %q", err, tt.Texts())
	}
	tt.Advance(time.Second)
	if !tt.HasText("Lobe UI is an open-source component library for AI apps.") {
		t.Errorf("accordion/basic: answer not shown after clicking the question; texts %q", tt.Texts())
	}
}

func TestDataEntryDemos(t *testing.T) {
	for _, demo := range []string{
		"radio/basic", "radio/vertical", "select/basic", "select/variants", "slider/basic",
		"search-bar/basic", "toggle-group/basic", "toggle-group/variants",
	} {
		if _, ok := registry.Get(demo); !ok {
			t.Errorf("%s not registered", demo)
		}
	}
	view, ok := registry.Get("select/basic")
	if !ok {
		return
	}
	tt := ui.NewTester(func(ctx *ui.Context) { preview(ctx, view) }, 720, 400)
	if err := tt.Click("Select a model"); err != nil {
		t.Fatalf("select/basic: %v; texts %q", err, tt.Texts())
	}
	tt.Advance(time.Second)
	if !tt.HasText("Claude Opus") {
		t.Errorf("select/basic: options not shown; texts %q", tt.Texts())
	}
}

func TestDataEntryBatchTwoDemos(t *testing.T) {
	for _, demo := range []string{
		"input-number/basic", "input-password/basic", "rate/basic", "rate/half", "copy-button/basic", "editable-text/basic",
	} {
		if _, ok := registry.Get(demo); !ok {
			t.Errorf("%s not registered", demo)
		}
	}
	view, ok := registry.Get("editable-text/basic")
	if !ok {
		return
	}
	tt := ui.NewTester(func(ctx *ui.Context) { preview(ctx, view) }, 720, 400)
	info, ok := tt.AccessByRole(ui.RoleButton, "Edit")
	if !ok {
		t.Fatalf("editable-text/basic: no Edit button; texts %q", tt.Texts())
	}
	tt.ClickAt(info.Bounds.X+info.Bounds.W/2, info.Bounds.Y+info.Bounds.H/2)
	tt.Frame()
	if _, ok := tt.AccessByRole(ui.RoleTextField, ""); !ok {
		t.Error("editable-text/basic: no field after clicking Edit")
	}
}

func TestDataEntryBatchThreeDemos(t *testing.T) {
	for _, demo := range []string{"auto-complete/basic", "hotkey-input/basic", "color-swatches/basic", "input-otp/basic", "color-picker/basic", "color-picker/triggers"} {
		if _, ok := registry.Get(demo); !ok {
			t.Errorf("%s not registered", demo)
		}
	}
}
