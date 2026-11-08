package divider

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func Vertical(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Row(c).Gap(10).Children(func() {
		for i, label := range []string{"Edit", "View", "Help"} {
			if i > 0 {
				layout.Divider(c.Key(i), layout.DividerProps{Vertical: true})
			}
			ui.Text(c.Key(label), label).TextColor(tk.TextSecondary)
		}
	})
}
