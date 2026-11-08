package colorpicker

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var (
	text      = ui.Hex("#13c2c2")
	highlight = ui.RGBA(250, 219, 20, 0.5)
	small     = ui.Hex("#eb2f96")
)

func Triggers(c *ui.Context) {
	ui.Row(c).Gap(16).Height(420).AlignItems(ui.Start).Children(func() {
		choice.ColorPicker(c, &text, choice.ColorPickerProps{Trigger: choice.ColorTriggerField}).Label("Text color")
		choice.ColorPicker(c, &highlight, choice.ColorPickerProps{Trigger: choice.ColorTriggerField, Alpha: true}).Label("Highlight")
		choice.ColorPicker(c, &small, choice.ColorPickerProps{Trigger: choice.ColorTriggerChip, Size: theme.Small}).Label("Tag color")
	})
}
