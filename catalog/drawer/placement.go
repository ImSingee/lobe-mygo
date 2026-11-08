package drawer

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var placementOpen = map[feedback.Placement]*bool{
	feedback.Left: new(bool), feedback.Right: new(bool), feedback.Top: new(bool), feedback.Bottom: new(bool),
}

func Placement(c *ui.Context) {
	tk := theme.TokensOf(c)
	sides := []struct {
		label     string
		placement feedback.Placement
	}{{"Left", feedback.Left}, {"Right", feedback.Right}, {"Top", feedback.Top}, {"Bottom", feedback.Bottom}}
	ui.Row(c).Gap(12).Children(func() {
		for _, s := range sides {
			if action.Button(c.Key(s.label), s.label, action.ButtonProps{}).Clicked() {
				*placementOpen[s.placement] = true
			}
		}
	})
	for _, s := range sides {
		feedback.Drawer(c.Key("drawer"+s.label), placementOpen[s.placement], feedback.DrawerProps{Title: s.label + " drawer", Placement: s.placement, Size: 260}, func() {
			ui.Text(c, "Slides in from the "+s.label+" edge.").TextColor(tk.TextSecondary)
		})
	}
}
