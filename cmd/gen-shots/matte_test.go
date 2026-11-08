package main

import (
	"image"
	"image/color"
	"testing"
)

func over(c color.NRGBA, bg uint8) color.RGBA {
	a := float64(c.A) / 255
	mix := func(v uint8) uint8 { return uint8(float64(v)*a + float64(bg)*(1-a) + 0.5) }
	return color.RGBA{mix(c.R), mix(c.G), mix(c.B), 255}
}

func TestUnmatteRecoversColorAndAlpha(t *testing.T) {
	cases := []color.NRGBA{{200, 40, 40, 255}, {0, 0, 0, 128}, {30, 120, 220, 64}, {0, 0, 0, 0}, {255, 255, 255, 200}}
	black := image.NewRGBA(image.Rect(0, 0, len(cases), 1))
	white := image.NewRGBA(image.Rect(0, 0, len(cases), 1))
	for i, c := range cases {
		black.SetRGBA(i, 0, over(c, 0))
		white.SetRGBA(i, 0, over(c, 255))
	}
	got := unmatte(black, white)
	for i, want := range cases {
		g := got.NRGBAAt(i, 0)
		if d := int(g.A) - int(want.A); d < -2 || d > 2 {
			t.Errorf("%d: alpha %d, want %d", i, g.A, want.A)
			continue
		}
		if want.A < 16 {
			continue
		}
		for _, p := range [][2]uint8{{g.R, want.R}, {g.G, want.G}, {g.B, want.B}} {
			if d := int(p[0]) - int(p[1]); d < -4 || d > 4 {
				t.Errorf("%d: color %+v, want %+v", i, g, want)
				break
			}
		}
	}
}

func TestHalveAveragesEachBlock(t *testing.T) {
	src := image.NewRGBA(image.Rect(0, 0, 4, 2))
	for x, v := range []uint8{0, 200, 40, 40} {
		for y := range 2 {
			src.SetRGBA(x, y, color.RGBA{v, 255 - v, 0, 255})
		}
	}
	got := halve(src)
	if b := got.Bounds(); b.Dx() != 2 || b.Dy() != 1 {
		t.Fatalf("bounds %v, want 2x1", b)
	}
	for i, want := range []color.RGBA{{100, 155, 0, 255}, {40, 215, 0, 255}} {
		if g := got.RGBAAt(i, 0); g != want {
			t.Errorf("%d: %+v, want %+v", i, g, want)
		}
	}
}
