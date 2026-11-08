package switchdemo

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var wifi, bluetooth, airplane = true, false, true

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	rows := []struct {
		label string
		on    *bool
		props choice.SwitchProps
		off   bool
	}{
		{"Wi-Fi", &wifi, choice.SwitchProps{}, false},
		{"Bluetooth (small)", &bluetooth, choice.SwitchProps{Size: theme.Small}, false},
		{"Airplane mode (disabled)", &airplane, choice.SwitchProps{}, true},
	}
	ui.Column(c).Gap(12).Width(260).Children(func() {
		for _, r := range rows {
			ui.Row(c.Key(r.label)).Children(func() {
				ui.Text(c, r.label).Grow(1).TextColor(tk.TextSecondary)
				choice.Switch(c, r.on, r.props).Label(r.label).Disabled(r.off)
			})
		}
	})
}
