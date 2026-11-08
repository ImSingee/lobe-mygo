package grid

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	layout.Grid(c, layout.GridProps{}).Children(func() {
		for i := range 6 {
			ui.Column(c.Key(i)).Gap(4).Padding(16).Radius(tk.BorderRadiusLG).Background(tk.FillQuaternary).Children(func() {
				ui.Text(c, fmt.Sprintf("Agent %d", i+1)).FontWeight(600)
				ui.Text(c, "Drafts replies and keeps notes.").TextColor(tk.TextSecondary)
			})
		}
	})
}
