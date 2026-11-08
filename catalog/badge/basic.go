package badge

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Row(c).Gap(32).AlignItems(ui.Center).Children(func() {
		square := func() { ui.Box(c).Size(40, 40).Radius(8).Background(tk.FillSecondary) }
		feedback.Badge(c, feedback.BadgeProps{Count: 5}, square)
		feedback.Badge(c.Key("over"), feedback.BadgeProps{Count: 120}, square)
		feedback.Badge(c.Key("dot"), feedback.BadgeProps{Dot: true}, square)
		feedback.Badge(c.Key("small"), feedback.BadgeProps{Count: 3, Small: true, Color: tk.Success}, square)
		feedback.Badge(c.Key("avatar"), feedback.BadgeProps{Count: 12}, func() { display.Avatar(c, "Ada", display.AvatarProps{Shape: display.AvatarCircle, Size: 40}) })
	})
}
