package spin

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var loading = true

func Content(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Column(c).Gap(12).Width(400).Children(func() {
		choice.Switch(c, &loading, choice.SwitchProps{})
		ui.Column(c).Padding(16).Radius(8).Border(1, tk.BorderSecondary).BorderStyle(ui.BorderDashed).Children(func() {
			feedback.Spin(c, feedback.SpinProps{Tip: "Loading…", Spinning: &loading}, func() {
				ui.Column(c).Height(96).Children(func() { ui.Text(c, "content") })
			})
		})
	})
}
