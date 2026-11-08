package main

import (
	"flag"
	"image/png"
	"log"
	"os"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
)

var variants = []struct {
	name    string
	variant lobe.ButtonVariant
}{
	{"Default", lobe.Default},
	{"Primary", lobe.Primary},
	{"Dashed", lobe.Dashed},
	{"Filled", lobe.Filled},
	{"Text", lobe.Text},
	{"Link", lobe.Link},
}

var sizes = []lobe.ButtonSize{lobe.Small, lobe.Middle, lobe.Large}

func view(c *ui.Context) {
	lobe.Apply(c)
	ui.Row(c).Fill().AlignItems(ui.Stretch).Children(func() {
		for _, tk := range []*lobe.Tokens{lobe.Light(), lobe.Dark()} {
			ui.Themed(c, tk.Theme()).Grow(1).Padding(24).Gap(16).Children(func() {
				for _, size := range sizes {
					ui.Row(c).Gap(8).Children(func() {
						for _, v := range variants {
							lobe.Button(c, v.name, lobe.ButtonProps{Variant: v.variant, Size: size})
						}
					})
				}
			})
		}
	})
}

func main() {
	snapshot := flag.String("snapshot", "", "write a PNG of the gallery and exit")
	flag.Parse()
	if *snapshot != "" {
		f, err := os.Create(*snapshot)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := png.Encode(f, ui.Render(view, 1120, 200, 2)); err != nil {
			log.Fatal(err)
		}
		return
	}
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{Title: "lobe-mygo gallery", Width: 1120, Height: 200, Content: ui.View(view)})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
