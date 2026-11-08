package avatar

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Fallback(c *ui.Context) {
	ui.Row(c).Gap(12).Children(func() {
		display.Avatar(c, "Lobe", display.AvatarProps{Fallback: display.AvatarText})
		display.Avatar(c, "AI", display.AvatarProps{Fallback: display.AvatarText, Background: ui.Hex("#7858f9")})
		display.Avatar(c, "🤯", display.AvatarProps{})
		display.Avatar(c, "🦊", display.AvatarProps{Background: ui.Hex("#fde68a")})
		display.Avatar(c, "Lobe", display.AvatarProps{Fallback: display.AvatarText, Shape: display.AvatarCircle})
	})
}
