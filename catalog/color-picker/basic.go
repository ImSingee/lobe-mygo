package colorpicker

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
)

var (
	brand   = ui.Hex("#1677ff")
	overlay = ui.RGBA(114, 46, 209, 0.6)
	presets = []ui.Color{
		ui.Hex("#f5222d"), ui.Hex("#fa541c"), ui.Hex("#fa8c16"), ui.Hex("#fadb14"), ui.Hex("#52c41a"),
		ui.Hex("#13c2c2"), ui.Hex("#1677ff"), ui.Hex("#722ed1"), ui.Hex("#eb2f96"),
	}
)

func Basic(c *ui.Context) {
	ui.Row(c).Gap(16).Height(420).AlignItems(ui.Start).Children(func() {
		choice.ColorPicker(c, &brand, choice.ColorPickerProps{Presets: presets}).Label("Brand color")
		choice.ColorPicker(c, &overlay, choice.ColorPickerProps{Alpha: true, Presets: presets}).Label("Overlay color")
	})
}
