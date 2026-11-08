package actionicon

import (
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
	"github.com/lobehub/lobe-mygo/lucide"
)

func Variants(c *ui.Context) {
	ui.Row(c).Gap(12).Children(func() {
		lobe.ActionIcon(c, lucide.Settings(), lobe.ActionIconProps{Title: "Borderless"})
		lobe.ActionIcon(c, lucide.Settings(), lobe.ActionIconProps{Title: "Filled", Variant: lobe.ActionIconFilled})
		lobe.ActionIcon(c, lucide.Settings(), lobe.ActionIconProps{Title: "Outlined", Variant: lobe.ActionIconOutlined})
	})
}
