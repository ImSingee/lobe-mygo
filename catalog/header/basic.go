package header

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func logo(c *ui.Context) {
	ui.Text(c, "LobeHub").FontSize(18).FontWeight(700)
}

func actions(c *ui.Context) {
	action.ActionIcon(c, lucide.Search(), action.ActionIconProps{Title: "Search"})
	action.ActionIcon(c, lucide.Moon(), action.ActionIconProps{Title: "Theme"})
}

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Column(c).AlignSelf(ui.Stretch).Gap(24).Children(func() {
		layout.Header(c, layout.HeaderProps{
			Logo: func() { logo(c) },
			Nav: func() {
				ui.Row(c).Gap(24).Children(func() {
					for _, l := range []string{"Components", "Docs", "Changelog"} {
						ui.Text(c, l).TextColor(tk.TextSecondary)
					}
				})
			},
			Actions: func() { actions(c) },
		})
		layout.Header(c, layout.HeaderProps{
			Mobile:  true,
			Nav:     func() { action.ActionIcon(c, lucide.Menu(), action.ActionIconProps{Title: "Menu"}) },
			Logo:    func() { ui.Row(c).Grow(1).Justify(ui.Center).Children(func() { logo(c) }) },
			Actions: func() { actions(c) },
		})
	})
}
