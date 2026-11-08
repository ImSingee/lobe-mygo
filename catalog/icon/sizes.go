package icon

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

func Sizes(c *ui.Context) {
	ui.Row(c).Gap(24).AlignItems(ui.Center).Children(func() {
		for _, size := range []display.IconSize{display.IconSmall, display.IconMiddle, display.IconLarge} {
			ui.Row(c).Gap(8).Children(func() {
				display.Icon(c, lucide.Smile(), display.IconProps{Size: size})
				display.Icon(c, lucide.Settings(), display.IconProps{Size: size})
				display.Icon(c, lucide.Search(), display.IconProps{Size: size})
			})
		}
	})
	ui.Row(c).Gap(6).FontSize(16).Children(func() {
		display.Icon(c, lucide.Sparkles(), display.IconProps{})
		ui.Text(c, "An icon without a size follows the text around it")
	})
}
