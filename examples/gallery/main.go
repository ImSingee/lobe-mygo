package main

import (
	"flag"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/theme"
	"image/png"
	"log"
	"os"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

var variants = []struct {
	name    string
	variant action.ButtonVariant
}{
	{"Default", action.Default},
	{"Primary", action.Primary},
	{"Dashed", action.Dashed},
	{"Filled", action.Filled},
	{"Text", action.Text},
	{"Link", action.Link},
}

var sizes = []theme.Size{theme.Small, theme.Middle, theme.Large}

func view(c *ui.Context) {
	theme.Apply(c)
	ui.Row(c).Fill().AlignItems(ui.Stretch).Children(func() {
		for _, tk := range []*theme.Tokens{theme.Light(), theme.Dark()} {
			ui.Themed(c, tk.Theme()).Grow(1).Padding(24).Gap(16).Children(func() {
				for _, size := range sizes {
					ui.Row(c).Gap(8).Children(func() {
						for _, v := range variants {
							action.Button(c, v.name, action.ButtonProps{Variant: v.variant, Size: size})
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
