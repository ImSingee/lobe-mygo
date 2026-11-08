package tag

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

var presets = []display.TagColor{
	display.TagRed, display.TagVolcano, display.TagOrange, display.TagGold, display.TagYellow, display.TagLime, display.TagGreen,
	display.TagCyan, display.TagBlue, display.TagGeekblue, display.TagPurple, display.TagMagenta, display.TagGray,
}

var statuses = []display.TagColor{display.TagSuccess, display.TagProcessing, display.TagWarning, display.TagError}

func Colors(c *ui.Context) {
	ui.Row(c).Wrap().Gap(8).Justify(ui.Center).Children(func() {
		for _, color := range presets {
			display.Tag(c, string(color), display.TagProps{Color: color})
		}
	})
	ui.Row(c).Gap(8).Children(func() {
		for _, color := range statuses {
			display.Tag(c, string(color), display.TagProps{Color: color})
		}
		display.Tag(c, "#7c3aed", display.TagProps{Color: display.TagColor("#7c3aed")})
	})
}
