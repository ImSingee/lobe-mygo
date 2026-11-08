package drawer

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var (
	settingsOpen     bool
	mentions, digest = true, false
)

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	if action.Button(c, "Open drawer", action.ButtonProps{Variant: action.Primary}).Clicked() {
		settingsOpen = true
	}
	feedback.Drawer(c, &settingsOpen, feedback.DrawerProps{
		Title: "Notification settings",
		Footer: func() {
			if action.Button(c, "Cancel", action.ButtonProps{}).Clicked() {
				settingsOpen = false
			}
			if action.Button(c, "Save", action.ButtonProps{Variant: action.Primary}).Clicked() {
				settingsOpen = false
			}
		},
	}, func() {
		ui.Column(c).Gap(16).Children(func() {
			for _, row := range []struct {
				label string
				on    *bool
			}{{"Mentions and replies", &mentions}, {"Daily digest", &digest}} {
				ui.Row(c.Key(row.label)).Children(func() {
					ui.Text(c, row.label).Grow(1).TextColor(tk.TextSecondary)
					choice.Switch(c, row.on, choice.SwitchProps{}).Label(row.label)
				})
			}
		})
	})
}
