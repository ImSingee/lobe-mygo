package alert

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
)

var dismissed bool

func Description(c *ui.Context) {
	ui.Column(c).Gap(12).Width(480).Children(func() {
		feedback.Alert(c, "Scheduled maintenance", feedback.AlertProps{
			Type:        feedback.Warning,
			Description: "The service is read only on Sunday from 02:00 to 03:00 UTC while the database moves to a new region.",
			Closed:      &dismissed,
		})
		if dismissed {
			ui.Text(c, "Dismissed. Reload the demo to see it again.")
		}
	})
}
