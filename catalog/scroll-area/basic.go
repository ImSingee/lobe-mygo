package scrollarea

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	layout.ScrollArea(c, layout.ScrollAreaProps{Fade: true}, func() {
		ui.Column(c).Gap(2).Padding(4).Children(func() {
			for i := range 40 {
				ui.Row(c.Key(i)).Padding(8, 12).Radius(tk.BorderRadiusSM).Background(tk.FillQuaternary).Children(func() {
					ui.Text(c, fmt.Sprintf("Message %d", i+1))
				})
			}
		})
	}).Width(360).Height(240)
}
