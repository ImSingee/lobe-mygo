package tabs

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/nav"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var verticalTab int

func Vertical(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Column(c).Width(420).Children(func() {
		nav.Tabs(c, &verticalTab, nav.TabsProps{Items: items[:3], Vertical: true}, func(i int) {
			ui.Text(c, panels[i]).TextColor(tk.TextSecondary)
		})
	})
}
