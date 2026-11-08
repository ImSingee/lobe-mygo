package alert

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(12).Width(480).Children(func() {
		feedback.Alert(c, "A new version is ready to install.", feedback.AlertProps{})
		feedback.Alert(c, "Your changes were saved.", feedback.AlertProps{Type: feedback.Success})
		feedback.Alert(c, "Your plan renews in 3 days.", feedback.AlertProps{Type: feedback.Warning})
		feedback.Alert(c, "The upload failed. Check your connection.", feedback.AlertProps{Type: feedback.Error})
		feedback.Alert(c, "This workspace is read only.", feedback.AlertProps{Type: feedback.Secondary})
	})
}
