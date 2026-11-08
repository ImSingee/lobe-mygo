package layoutdemo

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var pages = []display.ListItem{
	{Key: "intro", Label: "Introduction", Icon: lucide.House()},
	{Key: "install", Label: "Installation", Icon: lucide.BookOpen()},
	{Key: "theme", Label: "Theming", Icon: lucide.Moon()},
	{Key: "faq", Label: "FAQ", Icon: lucide.CircleHelp()},
}

var sections = []string{"Overview", "Quick start", "Components", "Customizing", "Next steps"}

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Box(c).Absolute().Left(0).Top(0).Right(0).Bottom(0)
	active := ui.Local(root, "page", func() string { return "intro" })
	root.Children(func() {
		layout.Layout(c, layout.LayoutProps{
			Header: func() {
				layout.Header(c, layout.HeaderProps{
					Logo: func() { ui.Text(c, "LobeHub").FontSize(18).FontWeight(700) },
					Nav: func() {
						ui.Row(c).Gap(24).Children(func() {
							ui.Text(c, "Docs")
							ui.Text(c, "Components").TextColor(tk.TextSecondary)
						})
					},
					Actions: func() {
						action.ActionIcon(c, lucide.Search(), action.ActionIconProps{Title: "Search"})
					},
				})
			},
			Sidebar: func() {
				ui.Column(c).Padding(12).Children(func() {
					if k := display.List(c, display.ListProps{Items: pages, Active: *active}); k != "" {
						*active = k
					}
				})
			},
			Toc: func() {
				ui.Column(c).Padding(24, 16).Gap(10).Children(func() {
					ui.Text(c, "On this page").FontSize(12).FontWeight(600).TextColor(tk.TextTertiary)
					for _, s := range sections {
						ui.Text(c.Key(s), s).FontSize(13).TextColor(tk.TextSecondary)
					}
				})
			},
			TocWidth: 200,
			Footer: func() {
				ui.Row(c).Padding(24).Justify(ui.Center).BorderWidth(1, 0, 0, 0).BorderColor(tk.BorderSecondary).Children(func() {
					ui.Text(c, "Copyright © 2026 LobeHub").FontSize(12).TextColor(tk.TextTertiary)
				})
			},
		}, func() {
			ui.Column(c).Padding(32, 40).Gap(16).Children(func() {
				ui.Text(c, "Page: "+*active).FontSize(28).FontWeight(700)
				for i, s := range sections {
					ui.Column(c.Key(i)).Gap(8).Children(func() {
						ui.Text(c, s).FontSize(18).FontWeight(600)
						ui.Text(c, fmt.Sprintf("The content scrolls under the header while the sidebar and the table of contents stay. Paragraph %d of the page.", i+1)).
							TextColor(tk.TextSecondary).LineHeight(1.7)
						ui.Box(c).Height(140).Radius(12).Background(tk.FillQuaternary)
					})
				}
			})
		})
	})
}
