package slider

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
)

var (
	temperature = 0.7
	maxTokens   = 4096.0
	locked      = 30.0
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(20).Width(360).Children(func() {
		row := func(label, value string, slider func()) {
			ui.Column(c.Key(label)).Gap(6).Children(func() {
				ui.Row(c).Justify(ui.SpaceBetween).FontSize(12).Children(func() {
					ui.Text(c, label)
					ui.Text(c, value)
				})
				slider()
			})
		}
		row("Temperature", fmt.Sprintf("%.1f", temperature), func() {
			choice.Slider(c, &temperature, choice.SliderProps{Max: 2, Step: 0.1}).Label("Temperature")
		})
		row("Max tokens", fmt.Sprintf("%.0f", maxTokens), func() {
			choice.Slider(c, &maxTokens, choice.SliderProps{Min: 256, Max: 8192, Step: 256}).Label("Max tokens")
		})
		row("Disabled", fmt.Sprintf("%.0f", locked), func() {
			choice.Slider(c, &locked, choice.SliderProps{Disabled: true}).Label("Disabled")
		})
	})
}
