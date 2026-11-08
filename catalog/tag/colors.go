package tag

import (
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
)

var presets = []lobe.TagColor{
	lobe.TagRed, lobe.TagVolcano, lobe.TagOrange, lobe.TagGold, lobe.TagYellow, lobe.TagLime, lobe.TagGreen,
	lobe.TagCyan, lobe.TagBlue, lobe.TagGeekblue, lobe.TagPurple, lobe.TagMagenta, lobe.TagGray,
}

var statuses = []lobe.TagColor{lobe.TagSuccess, lobe.TagProcessing, lobe.TagWarning, lobe.TagError}

func Colors(c *ui.Context) {
	ui.Row(c).Wrap().Gap(8).Justify(ui.Center).Children(func() {
		for _, color := range presets {
			lobe.Tag(c, string(color), lobe.TagProps{Color: color})
		}
	})
	ui.Row(c).Gap(8).Children(func() {
		for _, color := range statuses {
			lobe.Tag(c, string(color), lobe.TagProps{Color: color})
		}
		lobe.Tag(c, "#7c3aed", lobe.TagProps{Color: lobe.TagColor("#7c3aed")})
	})
}
