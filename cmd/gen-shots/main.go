// Command gen-shots renders the docs site's images from real components: the home page's bento
// tiles with transparent backgrounds and the live showcase's poster, in light and dark, the
// social preview and the README banner.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"slices"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/catalog/home"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

const scale = 2

func main() {
	out := flag.String("out", "public/home", "directory the images go to")
	data := flag.String("snippets", "docs/home/snippets.json", "file the code samples go to")
	banner := flag.String("banner", ".github/banner.png", "file the README banner goes to")
	flag.Parse()
	if err := os.MkdirAll(*out, 0o755); err != nil {
		log.Fatal(err)
	}
	for _, mode := range []string{"light", "dark"} {
		dark := mode == "dark"
		tiles := slices.Clone(home.Tiles)
		for _, sn := range home.Snippets {
			tiles = append(tiles, sn.Tile)
		}
		for _, t := range tiles {
			write(filepath.Join(*out, fmt.Sprintf("%s-%s.png", t.Name, mode)), tile(t, dark))
		}
		write(filepath.Join(*out, "showcase-"+mode+".png"), shot(home.Showcase, 1100, 640, dark, nil, 20*time.Second))
	}
	writeSnippets(*data)
	write(filepath.Join(filepath.Dir(*out), "og.png"), halve(shotAt(home.OG, 1280, 640, 2, true, nil, time.Second)))
	black := halve(shotAt(home.Banner, 1474, 602, 2, true, &color.RGBA{0, 0, 0, 255}, time.Second))
	white := halve(shotAt(home.Banner, 1474, 602, 2, true, &color.RGBA{255, 255, 255, 255}, time.Second))
	write(*banner, unmatte(black, white))
}

type snippet struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Code  string `json:"code"`
	Image string `json:"image"`
}

func writeSnippets(path string) {
	list := make([]snippet, len(home.Snippets))
	for i, sn := range home.Snippets {
		list[i] = snippet{Key: sn.Key, Label: sn.Label, Code: sn.Code, Image: sn.Tile.Name}
	}
	b, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(path, append(b, '\n'), 0o644); err != nil {
		log.Fatal(err)
	}
	log.Println("wrote", path)
}

func write(path string, img image.Image) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
	log.Println("wrote", path)
}

func tile(t home.Tile, dark bool) *image.NRGBA {
	black := shot(t.View, t.Width, t.Height, dark, &color.RGBA{0, 0, 0, 255}, time.Second)
	white := shot(t.View, t.Width, t.Height, dark, &color.RGBA{255, 255, 255, 255}, time.Second)
	return unmatte(black, white)
}

func shot(view func(c *ui.Context), w, h float32, dark bool, bg *color.RGBA, settle time.Duration) *image.RGBA {
	return shotAt(view, w, h, scale, dark, bg, settle)
}

func shotAt(view func(c *ui.Context), w, h, scale float32, dark bool, bg *color.RGBA, settle time.Duration) *image.RGBA {
	tt := ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		frame := ui.Column(c).Fill().Center()
		if bg != nil {
			frame.Background(ui.RGB(bg.R, bg.G, bg.B))
		}
		frame.Children(func() { view(c) })
	}, int(w), int(h))
	tt.SetScale(scale)
	tt.SetDark(dark)
	tt.Frame()
	tt.Advance(settle)
	tt.Frame()
	return tt.Image()
}

// unmatte recovers the colors and coverage of what was drawn over black and over white: the
// difference between the two is how much of the background shows through.
func unmatte(black, white *image.RGBA) *image.NRGBA {
	b := black.Bounds()
	out := image.NewNRGBA(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			k, w := black.RGBAAt(x, y), white.RGBAAt(x, y)
			through := max(int(w.R)-int(k.R), int(w.G)-int(k.G), int(w.B)-int(k.B), 0)
			a := 255 - min(through, 255)
			if a == 0 {
				continue
			}
			un := func(v uint8) uint8 { return uint8(min(int(v)*255/a, 255)) }
			out.SetNRGBA(x, y, color.NRGBA{un(k.R), un(k.G), un(k.B), uint8(a)})
		}
	}
	return out
}

func halve(src *image.RGBA) *image.RGBA {
	b := src.Bounds()
	out := image.NewRGBA(image.Rect(0, 0, b.Dx()/2, b.Dy()/2))
	for y := range b.Dy() / 2 {
		for x := range b.Dx() / 2 {
			var sum [4]int
			for _, p := range [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}} {
				c := src.RGBAAt(b.Min.X+2*x+p[0], b.Min.Y+2*y+p[1])
				sum[0] += int(c.R)
				sum[1] += int(c.G)
				sum[2] += int(c.B)
				sum[3] += int(c.A)
			}
			out.SetRGBA(x, y, color.RGBA{uint8((sum[0] + 2) / 4), uint8((sum[1] + 2) / 4), uint8((sum[2] + 2) / 4), uint8((sum[3] + 2) / 4)})
		}
	}
	return out
}
