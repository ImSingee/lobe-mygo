package main

import (
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/theme"
	"log"
	"os"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	_ "github.com/lobehub/lobe-mygo/catalog/action-icon"
	_ "github.com/lobehub/lobe-mygo/catalog/button"
	_ "github.com/lobehub/lobe-mygo/catalog/checkbox"
	_ "github.com/lobehub/lobe-mygo/catalog/icon"
	_ "github.com/lobehub/lobe-mygo/catalog/input"
	_ "github.com/lobehub/lobe-mygo/catalog/motion"
	"github.com/lobehub/lobe-mygo/catalog/registry"
	_ "github.com/lobehub/lobe-mygo/catalog/segmented"
	_ "github.com/lobehub/lobe-mygo/catalog/switch"
	_ "github.com/lobehub/lobe-mygo/catalog/tag"
	_ "github.com/lobehub/lobe-mygo/catalog/text-area"
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
	theme.Apply(c)
	ui.Column(c).Fill().Center().Gap(12).Padding(20).
		Background(theme.TokensOf(c).BgContainer).
		Children(func() { view(c) })
}

var chosen string

func index(c *ui.Context) {
	theme.Apply(c)
	names := registry.Names()
	if chosen == "" && len(names) > 0 {
		chosen = names[0]
	}
	tk := theme.TokensOf(c)
	ui.Row(c).Fill().AlignItems(ui.Stretch).Children(func() {
		ui.Scroll(c).Width(200).Background(tk.BgLayout).Children(func() {
			ui.Column(c).Padding(8).Gap(2).Children(func() {
				for _, name := range names {
					variant := action.Text
					if name == chosen {
						variant = action.Filled
					}
					if action.Button(c.Key(name), name, action.ButtonProps{Variant: variant}).AlignSelf(ui.Stretch).Clicked() {
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
