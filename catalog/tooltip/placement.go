package tooltip

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/overlay"
)

func Placement(c *ui.Context) {
	sides := []struct {
		label     string
		placement overlay.Placement
	}{{"Top", overlay.Top}, {"Bottom", overlay.Bottom}, {"Left", overlay.Left}, {"Right", overlay.Right}}
	ui.Row(c).Gap(12).Children(func() {
		for _, s := range sides {
			b := action.Button(c.Key(s.label), s.label, action.ButtonProps{})
			overlay.Tooltip(c, b, "Shows on the "+s.label, overlay.TooltipProps{Placement: s.placement})
		}
	})
}
