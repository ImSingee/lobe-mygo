package progress

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
)

func Variants(c *ui.Context) {
	ui.Column(c).Gap(16).Width(360).Children(func() {
		feedback.Progress(c, 65, feedback.ProgressProps{Variant: feedback.ProgressSegments, ShowInfo: true})
		feedback.Progress(c, 65, feedback.ProgressProps{Variant: feedback.ProgressSegments, Segments: 10, ShowInfo: true})
		feedback.Progress(c, 65, feedback.ProgressProps{Variant: feedback.ProgressInset, ShowInfo: true})
	})
}
