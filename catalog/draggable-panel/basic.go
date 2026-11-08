package draggablepanel

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var (
	sideSize     float32 = 260
	sideExpanded         = true
)

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	frame(c, tk, func() {
		ui.Column(c).Grow(1).Padding(16).Gap(6).Children(func() {
			ui.Text(c, "Conversation").FontWeight(600)
			ui.Text(c, "Drag the seam to resize the panel, or use the button on it to collapse it.").TextColor(tk.TextSecondary)
		})
		layout.DraggablePanel(c, &sideSize, &sideExpanded, layout.DraggablePanelProps{MinSize: 200, MaxSize: 420}, func() {
			ui.Column(c).Padding(16).Gap(6).Children(func() {
				ui.Text(c, "Details").FontWeight(600)
				ui.Text(c, "Participants, files and pinned messages.").TextColor(tk.TextSecondary)
			})
		})
	})
}

func frame(c *ui.Context, tk *theme.Tokens, children func()) {
	ui.Row(c).AlignSelf(ui.Stretch).Height(260).AlignItems(ui.Stretch).
		Radius(tk.BorderRadiusLG).Border(1, tk.FillBorder()).Clip().Children(children)
}
