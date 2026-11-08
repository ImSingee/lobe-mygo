package skeleton

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Basic(c *ui.Context) {
	ui.Row(c).Gap(16).Width(480).AlignItems(ui.Start).Children(func() {
		display.SkeletonAvatar(c, 0, display.AvatarCircle, display.SkeletonFade)
		ui.Column(c).Grow(1).Gap(8).Children(func() {
			display.Skeleton(c, display.SkeletonProps{Width: 160, Height: 20})
			display.SkeletonText(c, display.SkeletonTextProps{Rows: 3})
		})
	})
}
