package skeleton

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Animation(c *ui.Context) {
	ui.Column(c).Gap(20).Width(480).Children(func() {
		for _, a := range []struct {
			name string
			anim display.SkeletonAnimation
		}{{"Fade", display.SkeletonFade}, {"Sweep", display.SkeletonSweep}, {"Static", display.SkeletonStatic}} {
			ui.Column(c.Key(a.name)).Gap(8).Children(func() {
				ui.Text(c, a.name).FontSize(12)
				display.SkeletonText(c, display.SkeletonTextProps{Rows: 2, Animation: a.anim})
			})
		}
	})
}
