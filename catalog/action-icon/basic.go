package actionicon

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

var pinned bool

func Basic(c *ui.Context) {
	ui.Row(c).Gap(4).Children(func() {
		action.ActionIcon(c, lucide.Copy(), action.ActionIconProps{Title: "Copy"})
		action.ActionIcon(c, lucide.Share2(), action.ActionIconProps{Title: "Share"})
		if action.ActionIcon(c, lucide.Pin(), action.ActionIconProps{Title: "Pin", Active: pinned}).Clicked() {
			pinned = !pinned
		}
		action.ActionIcon(c, lucide.Trash(), action.ActionIconProps{Title: "Delete", Danger: true})
	})
}
