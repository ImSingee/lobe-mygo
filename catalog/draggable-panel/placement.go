package draggablepanel

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var (
	navSize     float32 = 220
	navExpanded         = true
)

func Placement(c *ui.Context) {
	tk := theme.TokensOf(c)
	frame(c, tk, func() {
		layout.DraggablePanel(c, &navSize, &navExpanded, layout.DraggablePanelProps{Placement: layout.Left, CollapseThreshold: 120}, func() {
			ui.Column(c).Padding(16).Gap(6).Children(func() {
				ui.Text(c, "Sessions").FontWeight(600)
				ui.Text(c, "Drag below 120 DIPs to collapse.").TextColor(tk.TextSecondary)
			})
		})
		ui.Column(c).Grow(1).Padding(16).Children(func() {
			ui.Text(c, "Main content").TextColor(tk.TextSecondary)
		})
	})
}
