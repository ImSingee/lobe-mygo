package colorswatches

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
)

var (
	accent  = ui.Hex("#1677ff")
	label   = ui.Hex("#52c41a")
	palette = []choice.ColorSwatch{
		{Color: ui.Hex("#f5222d"), Title: "Red"}, {Color: ui.Hex("#fa8c16"), Title: "Orange"},
		{Color: ui.Hex("#fadb14"), Title: "Yellow"}, {Color: ui.Hex("#52c41a"), Title: "Green"},
		{Color: ui.Hex("#13c2c2"), Title: "Cyan"}, {Color: ui.Hex("#1677ff"), Title: "Blue"},
		{Color: ui.Hex("#722ed1"), Title: "Purple"}, {Color: ui.Hex("#eb2f96"), Title: "Magenta"},
		{Color: ui.Transparent, Title: "None"},
	}
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Height(440).Children(func() {
		choice.ColorSwatches(c, &accent, choice.ColorSwatchesProps{Colors: palette, EnableColorPicker: true}).Label("Accent")
		choice.ColorSwatches(c, &label, choice.ColorSwatchesProps{Colors: palette, Square: true, Size: 28}).Label("Label")
	})
}
