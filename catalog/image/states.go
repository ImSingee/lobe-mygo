package image

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func States(c *ui.Context) {
	tk := theme.TokensOf(c)
	cell := func(label string, img func()) {
		ui.Column(c.Key(label)).Gap(8).AlignItems(ui.Center).Children(func() {
			img()
			ui.Text(c, label).FontSize(12).TextColor(tk.TextSecondary)
		})
	}
	ui.Row(c).Gap(16).AlignItems(ui.Start).Children(func() {
		cell("Contain", func() {
			display.Image(c, wave(), display.ImageProps{Size: 140, Contain: true, Variant: display.ImageOutlined})
		})
		cell("Actions", func() {
			display.Image(c, crest(), display.ImageProps{Size: 140, Actions: func() {
				ui.Row(c).Padding(6).Children(func() {
					action.ActionIcon(c, lucide.Download(), action.ActionIconProps{Size: theme.Small, Variant: action.ActionIconFilled, Title: "Download"})
				})
			}})
		})
		cell("Loading", func() { display.Image(c, nil, display.ImageProps{Size: 140, Loading: true}) })
		cell("Failed", func() { display.Image(c, nil, display.ImageProps{Size: 140, Failed: true}) })
	})
}
