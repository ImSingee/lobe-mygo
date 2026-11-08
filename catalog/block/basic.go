package block

import (
	"strconv"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var clicks int

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	card := func(title string, p layout.BlockProps) {
		b := layout.Block(c.Key(title), p).Size(160, 96).Padding(16).Gap(4)
		b.Children(func() {
			ui.Text(c, title).FontWeight(500)
			ui.Text(c, "Block surface").FontSize(12).TextColor(tk.TextSecondary)
		})
		if b.Clicked() {
			clicks++
		}
	}
	ui.Column(c).Gap(16).AlignItems(ui.Center).Children(func() {
		ui.Row(c).Gap(16).Children(func() {
			card("Filled", layout.BlockProps{})
			card("Outlined", layout.BlockProps{Variant: layout.BlockOutlined})
			card("Borderless", layout.BlockProps{Variant: layout.BlockBorderless})
		})
		ui.Row(c).Gap(16).Children(func() {
			card("Clickable", layout.BlockProps{Clickable: true})
			card("Outlined click", layout.BlockProps{Variant: layout.BlockOutlined, Clickable: true})
			card("Shadow", layout.BlockProps{Variant: layout.BlockOutlined, Shadow: true})
		})
		ui.Text(c, "Clicks: "+strconv.Itoa(clicks)).TextColor(tk.TextSecondary)
	})
}
