package steps

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/nav"
)

var current = 1

var setup = []nav.StepItem{
	{Title: "Account", Description: "Sign in"},
	{Title: "Profile", Description: "Name and avatar"},
	{Title: "Workspace", Description: "Invite your team"},
	{Title: "Done"},
}

func Basic(c *ui.Context) {
	ui.Column(c).Gap(20).Width(560).Children(func() {
		nav.Steps(c.Key("default"), nav.StepsProps{Items: setup, Current: current})
		nav.Steps(c.Key("dot"), nav.StepsProps{Items: setup, Current: current, Dot: true})
		ui.Row(c).Gap(8).Justify(ui.Center).Children(func() {
			if action.Button(c, "Back", action.ButtonProps{}).Clicked() && current > 0 {
				current--
			}
			if action.Button(c, "Next", action.ButtonProps{Variant: action.Primary}).Clicked() && current < len(setup) {
				current++
			}
		})
	})
}
