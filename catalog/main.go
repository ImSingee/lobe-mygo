package main

import (
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/theme"
	"log"
	"os"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	_ "github.com/lobehub/lobe-mygo/catalog/accordion"
	_ "github.com/lobehub/lobe-mygo/catalog/action-icon"
	_ "github.com/lobehub/lobe-mygo/catalog/action-icon-group"
	_ "github.com/lobehub/lobe-mygo/catalog/agent-avatar"
	_ "github.com/lobehub/lobe-mygo/catalog/alert"
	_ "github.com/lobehub/lobe-mygo/catalog/auto-complete"
	_ "github.com/lobehub/lobe-mygo/catalog/avatar"
	_ "github.com/lobehub/lobe-mygo/catalog/badge"
	_ "github.com/lobehub/lobe-mygo/catalog/block"
	_ "github.com/lobehub/lobe-mygo/catalog/breadcrumb"
	_ "github.com/lobehub/lobe-mygo/catalog/burger"
	_ "github.com/lobehub/lobe-mygo/catalog/button"
	_ "github.com/lobehub/lobe-mygo/catalog/checkbox"
	_ "github.com/lobehub/lobe-mygo/catalog/color-picker"
	_ "github.com/lobehub/lobe-mygo/catalog/color-swatches"
	_ "github.com/lobehub/lobe-mygo/catalog/context-menu"
	_ "github.com/lobehub/lobe-mygo/catalog/copy-button"
	_ "github.com/lobehub/lobe-mygo/catalog/date-picker"
	_ "github.com/lobehub/lobe-mygo/catalog/divider"
	_ "github.com/lobehub/lobe-mygo/catalog/draggable-panel"
	_ "github.com/lobehub/lobe-mygo/catalog/draggable-side-nav"
	_ "github.com/lobehub/lobe-mygo/catalog/drawer"
	_ "github.com/lobehub/lobe-mygo/catalog/dropdown-menu"
	_ "github.com/lobehub/lobe-mygo/catalog/editable-text"
	_ "github.com/lobehub/lobe-mygo/catalog/empty"
	_ "github.com/lobehub/lobe-mygo/catalog/floating-panel"
	_ "github.com/lobehub/lobe-mygo/catalog/floating-sheet"
	_ "github.com/lobehub/lobe-mygo/catalog/form"
	_ "github.com/lobehub/lobe-mygo/catalog/grid"
	_ "github.com/lobehub/lobe-mygo/catalog/group-avatar"
	_ "github.com/lobehub/lobe-mygo/catalog/guide-card"
	_ "github.com/lobehub/lobe-mygo/catalog/header"
	_ "github.com/lobehub/lobe-mygo/catalog/highlighter"
	_ "github.com/lobehub/lobe-mygo/catalog/home"
	_ "github.com/lobehub/lobe-mygo/catalog/hotkey"
	_ "github.com/lobehub/lobe-mygo/catalog/hotkey-input"
	_ "github.com/lobehub/lobe-mygo/catalog/icon"
	_ "github.com/lobehub/lobe-mygo/catalog/image"
	_ "github.com/lobehub/lobe-mygo/catalog/image-select"
	_ "github.com/lobehub/lobe-mygo/catalog/input"
	_ "github.com/lobehub/lobe-mygo/catalog/input-number"
	_ "github.com/lobehub/lobe-mygo/catalog/input-otp"
	_ "github.com/lobehub/lobe-mygo/catalog/input-password"
	_ "github.com/lobehub/lobe-mygo/catalog/layout"
	_ "github.com/lobehub/lobe-mygo/catalog/list"
	_ "github.com/lobehub/lobe-mygo/catalog/markdown"
	_ "github.com/lobehub/lobe-mygo/catalog/modal"
	_ "github.com/lobehub/lobe-mygo/catalog/motion"
	_ "github.com/lobehub/lobe-mygo/catalog/pagination"
	_ "github.com/lobehub/lobe-mygo/catalog/popover"
	_ "github.com/lobehub/lobe-mygo/catalog/progress"
	_ "github.com/lobehub/lobe-mygo/catalog/qr-code"
	_ "github.com/lobehub/lobe-mygo/catalog/radio"
	_ "github.com/lobehub/lobe-mygo/catalog/rate"
	"github.com/lobehub/lobe-mygo/catalog/registry"
	_ "github.com/lobehub/lobe-mygo/catalog/result"
	_ "github.com/lobehub/lobe-mygo/catalog/scroll-area"
	_ "github.com/lobehub/lobe-mygo/catalog/search-bar"
	_ "github.com/lobehub/lobe-mygo/catalog/segmented"
	_ "github.com/lobehub/lobe-mygo/catalog/select"
	_ "github.com/lobehub/lobe-mygo/catalog/side-nav"
	_ "github.com/lobehub/lobe-mygo/catalog/skeleton"
	_ "github.com/lobehub/lobe-mygo/catalog/slider"
	_ "github.com/lobehub/lobe-mygo/catalog/snippet"
	_ "github.com/lobehub/lobe-mygo/catalog/sortable-list"
	_ "github.com/lobehub/lobe-mygo/catalog/spin"
	_ "github.com/lobehub/lobe-mygo/catalog/steps"
	_ "github.com/lobehub/lobe-mygo/catalog/switch"
	_ "github.com/lobehub/lobe-mygo/catalog/tabs"
	_ "github.com/lobehub/lobe-mygo/catalog/tag"
	_ "github.com/lobehub/lobe-mygo/catalog/text"
	_ "github.com/lobehub/lobe-mygo/catalog/text-area"
	_ "github.com/lobehub/lobe-mygo/catalog/toast"
	_ "github.com/lobehub/lobe-mygo/catalog/toggle-group"
	_ "github.com/lobehub/lobe-mygo/catalog/tokens"
	_ "github.com/lobehub/lobe-mygo/catalog/tooltip"
	_ "github.com/lobehub/lobe-mygo/catalog/upload"
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
