package main

import (
	"log"
	"os"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
	_ "github.com/lobehub/lobe-mygo/catalog/action-icon"
	_ "github.com/lobehub/lobe-mygo/catalog/button"
	_ "github.com/lobehub/lobe-mygo/catalog/icon"
	"github.com/lobehub/lobe-mygo/catalog/registry"
	_ "github.com/lobehub/lobe-mygo/catalog/tag"
	_ "github.com/lobehub/lobe-mygo/catalog/tokens"
)

func main() {
	content := index
	if len(os.Args) > 1 && os.Args[1] != "" {
		view, ok := registry.Get(os.Args[1])
		if !ok {
			log.Fatalf("catalog: no demo named %q; have %v", os.Args[1], registry.Names())
		}
		content = func(c *ui.Context) { preview(c, view) }
	}
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{Title: "Lobe MyGo", Width: 720, Height: 480, Content: ui.View(content)})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}

func preview(c *ui.Context, view func(c *ui.Context)) {
	lobe.Apply(c)
	ui.Column(c).Fill().Center().Gap(12).Padding(20).
		Background(lobe.TokensOf(c).BgContainer).
		Children(func() { view(c) })
}

var chosen string

func index(c *ui.Context) {
	lobe.Apply(c)
	names := registry.Names()
	if chosen == "" && len(names) > 0 {
		chosen = names[0]
	}
	tk := lobe.TokensOf(c)
	ui.Row(c).Fill().AlignItems(ui.Stretch).Children(func() {
		ui.Scroll(c).Width(200).Background(tk.BgLayout).Children(func() {
			ui.Column(c).Padding(8).Gap(2).Children(func() {
				for _, name := range names {
					variant := lobe.Text
					if name == chosen {
						variant = lobe.Filled
					}
					if lobe.Button(c.Key(name), name, lobe.ButtonProps{Variant: variant}).AlignSelf(ui.Stretch).Clicked() {
						chosen = name
					}
				}
			})
		})
		ui.Column(c).Grow(1).Children(func() {
			if view, ok := registry.Get(chosen); ok {
				preview(c, view)
			}
		})
	})
}
