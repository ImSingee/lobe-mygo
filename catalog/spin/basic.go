package spin

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var (
	thirty     float32 = 30
	seventyTwo float32 = 72
	sizes              = []theme.Size{theme.Small, theme.Middle, theme.Large}
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Children(func() {
		ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
			for _, s := range sizes {
				feedback.Spin(c.Key(s), feedback.SpinProps{Size: s}, nil)
			}
		})
		ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
			feedback.Spin(c, feedback.SpinProps{Percent: &thirty}, nil)
			feedback.Spin(c, feedback.SpinProps{Percent: &seventyTwo, Size: theme.Large}, nil)
		})
		ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
			for _, s := range sizes {
				feedback.Spin(c.Key(s), feedback.SpinProps{Variant: feedback.SpinNetwork, Size: s}, nil)
			}
		})
	})
}
