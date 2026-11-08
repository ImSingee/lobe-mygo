package alert

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
)

func Variants(c *ui.Context) {
	ui.Column(c).Gap(12).Width(480).Children(func() {
		feedback.Alert(c, "Soft tints the surface with the type's color.", feedback.AlertProps{Type: feedback.Success})
		feedback.Alert(c, "Outlined draws a border instead.", feedback.AlertProps{Type: feedback.Success, Variant: feedback.Outlined})
		feedback.Alert(c, "Plain shows the icon and text alone.", feedback.AlertProps{Type: feedback.Success, Variant: feedback.Plain})
	})
}
