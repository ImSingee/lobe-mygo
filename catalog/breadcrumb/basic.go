package breadcrumb

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/nav"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var visited = "nothing yet"

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	visit := func(page string) func() { return func() { visited = page } }
	ui.Column(c).Gap(16).AlignItems(ui.Center).Children(func() {
		nav.Breadcrumb(c.Key("chevron"), nav.BreadcrumbProps{Items: []nav.BreadcrumbItem{
			{Title: "Home", OnClick: visit("Home")},
			{Title: "Agents", OnClick: visit("Agents")},
			{Title: "Translator"},
		}})
		nav.Breadcrumb(c.Key("slash"), nav.BreadcrumbProps{Separator: "/", Items: []nav.BreadcrumbItem{
			{Title: "Workspace", OnClick: visit("Workspace")},
			{Title: "Files", OnClick: visit("Files")},
			{Title: "Reports"},
			{Title: "2026"},
		}})
		ui.Text(c, "Last visited: "+visited).TextColor(tk.TextSecondary)
	})
}
