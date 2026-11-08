package avatar

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Sizes(c *ui.Context) {
	ui.Column(c).Gap(16).Children(func() {
		for _, shape := range []display.AvatarShape{display.AvatarSquare, display.AvatarCircle} {
			ui.Row(c.Key(shape)).Gap(12).AlignItems(ui.End).Children(func() {
				for _, size := range []float32{20, 28, 36, 48, 64} {
					display.Avatar(c.Key(size), "Grace Hopper", display.AvatarProps{Size: size, Shape: shape})
				}
			})
		}
	})
}
