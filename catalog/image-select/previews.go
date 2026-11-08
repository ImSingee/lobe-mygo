package imageselect

import (
	"image"
	"image/color"
	"sync"

	"github.com/egoist/mygo/ui"
)

type palette struct{ page, sidebar, card, line, accent color.RGBA }

var (
	light = palette{rgb(0xf8, 0xf8, 0xf8), rgb(0xee, 0xee, 0xee), rgb(0xff, 0xff, 0xff), rgb(0xd9, 0xd9, 0xd9), rgb(0x08, 0x08, 0x08)}
	dark  = palette{rgb(0x11, 0x11, 0x11), rgb(0x1b, 0x1b, 0x1b), rgb(0x26, 0x26, 0x26), rgb(0x3a, 0x3a, 0x3a), rgb(0xf0, 0xf0, 0xf0)}
)

func rgb(r, g, b uint8) color.RGBA { return color.RGBA{r, g, b, 255} }

const previewW, previewH = 288, 172

func fill(img *image.RGBA, x0, y0, x1, y1 int, c color.RGBA, split func(x, y int) bool) {
	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			if split == nil || split(x, y) {
				img.SetRGBA(x, y, c)
			}
		}
	}
}

func paint(img *image.RGBA, p palette, split func(x, y int) bool) {
	fill(img, 0, 0, previewW, previewH, p.page, split)
	fill(img, 0, 0, 64, previewH, p.sidebar, split)
	for i := range 4 {
		fill(img, 14, 22+i*20, 50, 30+i*20, p.line, split)
	}
	fill(img, 14, 22, 30, 30, p.accent, split)
	fill(img, 84, 22, 268, 74, p.card, split)
	fill(img, 98, 36, 200, 44, p.accent, split)
	fill(img, 98, 52, 240, 58, p.line, split)
	fill(img, 84, 88, 170, 152, p.card, split)
	fill(img, 182, 88, 268, 152, p.card, split)
	fill(img, 98, 102, 150, 108, p.line, split)
	fill(img, 196, 102, 248, 108, p.line, split)
}

func preview(draw func(img *image.RGBA)) func() ui.ImageSource {
	return sync.OnceValue(func() ui.ImageSource {
		img := image.NewRGBA(image.Rect(0, 0, previewW, previewH))
		draw(img)
		return ui.NewBitmap(img)
	})
}

var (
	lightPreview = preview(func(img *image.RGBA) { paint(img, light, nil) })
	darkPreview  = preview(func(img *image.RGBA) { paint(img, dark, nil) })
	autoPreview  = preview(func(img *image.RGBA) {
		paint(img, light, nil)
		paint(img, dark, func(x, y int) bool { return x*previewH > (previewH-y)*previewW })
	})
)
