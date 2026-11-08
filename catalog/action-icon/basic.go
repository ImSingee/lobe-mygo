package actionicon

import (
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
	"github.com/lobehub/lobe-mygo/lucide"
)

var pinned bool

func Basic(c *ui.Context) {
	ui.Row(c).Gap(4).Children(func() {
		lobe.ActionIcon(c, lucide.Copy(), lobe.ActionIconProps{Title: "Copy"})
		lobe.ActionIcon(c, lucide.Share2(), lobe.ActionIconProps{Title: "Share"})
		if lobe.ActionIcon(c, lucide.Pin(), lobe.ActionIconProps{Title: "Pin", Active: pinned}).Clicked() {
			pinned = !pinned
		}
		lobe.ActionIcon(c, lucide.Trash(), lobe.ActionIconProps{Title: "Delete", Danger: true})
	})
}
