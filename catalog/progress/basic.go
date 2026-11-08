package progress

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var pct float32 = 40

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Width(360).Children(func() {
		feedback.Progress(c, pct, feedback.ProgressProps{Label: "Uploading report.pdf", ShowInfo: true})
		feedback.Progress(c, pct, feedback.ProgressProps{Status: feedback.ProgressActive, ShowInfo: true, Label: "Indexing"})
		feedback.Progress(c, 100, feedback.ProgressProps{Status: feedback.ProgressSuccess, ShowInfo: true, Label: "Done"})
		feedback.Progress(c, 62, feedback.ProgressProps{Status: feedback.ProgressException, ShowInfo: true, Label: "Failed"})
		ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
			feedback.Progress(c, pct, feedback.ProgressProps{Type: feedback.ProgressCircle, Size: theme.Small})
			feedback.Progress(c, pct, feedback.ProgressProps{Type: feedback.ProgressCircle, ShowInfo: true})
			feedback.Progress(c, pct, feedback.ProgressProps{Type: feedback.ProgressCircle, Size: theme.Large, ShowInfo: true})
		})
		ui.Row(c).Gap(8).Children(func() {
			if action.Button(c, "−10", action.ButtonProps{}).Clicked() {
				pct = max(pct-10, 0)
			}
			if action.Button(c, "+10", action.ButtonProps{}).Clicked() {
				pct = min(pct+10, 100)
			}
		})
	})
}
