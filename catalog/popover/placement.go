package popover

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/overlay"
)

var placementOpen = map[overlay.Placement]*bool{
	overlay.Top: new(bool), overlay.Bottom: new(bool), overlay.Left: new(bool), overlay.Right: new(bool),
}

func Placement(c *ui.Context) {
	sides := []struct {
		label     string
		placement overlay.Placement
	}{{"Top", overlay.Top}, {"Bottom", overlay.Bottom}, {"Left", overlay.Left}, {"Right", overlay.Right}}
	ui.Row(c).Gap(12).Children(func() {
		for _, s := range sides {
			b := action.Button(c.Key(s.label), s.label, action.ButtonProps{})
			overlay.Popover(c, b, placementOpen[s.placement], overlay.PopoverProps{Placement: s.placement}, func() {
				ui.Text(c, "Opens on the "+s.label)
			})
		}
	})
}
