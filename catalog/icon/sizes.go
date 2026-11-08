package icon

import (
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
	"github.com/lobehub/lobe-mygo/lucide"
)

func Sizes(c *ui.Context) {
	ui.Row(c).Gap(24).AlignItems(ui.Center).Children(func() {
		for _, size := range []lobe.IconSize{lobe.IconSmall, lobe.IconMiddle, lobe.IconLarge} {
			ui.Row(c).Gap(8).Children(func() {
				lobe.Icon(c, lucide.Smile(), lobe.IconProps{Size: size})
				lobe.Icon(c, lucide.Settings(), lobe.IconProps{Size: size})
				lobe.Icon(c, lucide.Search(), lobe.IconProps{Size: size})
			})
		}
	})
	ui.Row(c).Gap(6).FontSize(16).Children(func() {
		lobe.Icon(c, lucide.Sparkles(), lobe.IconProps{})
		ui.Text(c, "An icon without a size follows the text around it")
	})
}
