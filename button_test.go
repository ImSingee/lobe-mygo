package lobe

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func over(dst, src ui.Color) ui.Color {
	a := float32(src.A) / 255
	ch := func(d, s uint8) uint8 { return uint8(float32(s)*a + float32(d)*(1-a) + 0.5) }
	return ui.RGB(ch(dst.R, src.R), ch(dst.G, src.G), ch(dst.B, src.B))
}

func near(a, b ui.Color) bool {
	d := func(x, y uint8) int {
		if x > y {
			return int(x - y)
		}
		return int(y - x)
	}
	return d(a.R, b.R) <= 2 && d(a.G, b.G) <= 2 && d(a.B, b.B) <= 2
}

type buttonHarness struct {
	t       *testing.T
	tt      *ui.Tester
	clicked bool
}

func newButtonHarness(t *testing.T, props ButtonProps) *buttonHarness {
	h := &buttonHarness{t: t}
	h.tt = ui.NewTester(func(c *ui.Context) {
		Apply(c)
		ui.Column(c).Fill().Center().Children(func() {
			if Button(c, "Save", props).Clicked() {
				h.clicked = true
			}
		})
	}, 320, 200)
	return h
}

func (h *buttonHarness) bounds() ui.Rect {
	h.t.Helper()
	info, ok := h.tt.AccessByRole(ui.RoleButton, "Save")
	if !ok {
		h.t.Fatalf("no button named Save; texts %q", h.tt.Texts())
	}
	return info.Bounds
}

func (h *buttonHarness) fill() ui.Color {
	h.t.Helper()
	b := h.bounds()
	img := h.tt.Image()
	scale := float32(img.Bounds().Dx()) / 320
	px := img.RGBAAt(int((b.X+5)*scale), int((b.Y+b.H/2)*scale))
	return ui.RGB(px.R, px.G, px.B)
}

func (h *buttonHarness) center() (float32, float32) {
	b := h.bounds()
	return b.X + b.W/2, b.Y + b.H/2
}

func TestPrimaryButtonPaintsPrimary(t *testing.T) {
	h := newButtonHarness(t, ButtonProps{Variant: Primary})
	if got, want := h.fill(), Light().Primary; !near(got, want) {
		t.Errorf("fill = %+v, want Primary %+v", got, want)
	}
}

func TestPrimaryButtonFollowsDarkAppearance(t *testing.T) {
	h := newButtonHarness(t, ButtonProps{Variant: Primary})
	h.tt.SetDark(true)
	h.tt.Advance(time.Second)
	if got, want := h.fill(), Dark().Primary; !near(got, want) {
		t.Errorf("dark fill = %+v, want Primary %+v", got, want)
	}
}

func TestFilledButtonWashesOnHoverAndPress(t *testing.T) {
	tk := Light()
	h := newButtonHarness(t, ButtonProps{Variant: Filled})
	if got, want := h.fill(), over(tk.BgLayout, tk.FillTertiary); !near(got, want) {
		t.Errorf("rest = %+v, want FillTertiary over BgLayout %+v", got, want)
	}

	x, y := h.center()
	h.tt.Move(x, y)
	h.tt.Advance(time.Second)
	if got, want := h.fill(), over(tk.BgLayout, tk.FillSecondary); !near(got, want) {
		t.Errorf("hover = %+v, want FillSecondary over BgLayout %+v", got, want)
	}

	h.tt.Press(x, y)
	h.tt.Advance(time.Second)
	if got, want := h.fill(), over(tk.BgLayout, tk.Fill); !near(got, want) {
		t.Errorf("pressed = %+v, want Fill over BgLayout %+v", got, want)
	}
}

func TestButtonHeightsMatchLobeSizes(t *testing.T) {
	for _, c := range []struct {
		size   Size
		height float32
	}{{Small, 24}, {Middle, 32}, {Large, 40}} {
		h := newButtonHarness(t, ButtonProps{Size: c.size})
		if got := h.bounds().H; got != c.height {
			t.Errorf("size %v height = %v, want %v", c.size, got, c.height)
		}
	}
}

func TestButtonReportsClicked(t *testing.T) {
	h := newButtonHarness(t, ButtonProps{})
	if err := h.tt.Click("Save"); err != nil {
		t.Fatal(err)
	}
	if !h.clicked {
		t.Error("Clicked() did not report the click")
	}
}
