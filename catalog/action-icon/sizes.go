package actionicon

import (
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
	"github.com/lobehub/lobe-mygo/lucide"
)

func Sizes(c *ui.Context) {
	ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
		for _, size := range []lobe.Size{lobe.Small, lobe.Middle, lobe.Large} {
			lobe.ActionIcon(c, lucide.Bell(), lobe.ActionIconProps{Title: "Notifications", Size: size, Variant: lobe.ActionIconFilled})
		}
	})
}
