package actionicon

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

func Variants(c *ui.Context) {
	ui.Row(c).Gap(12).Children(func() {
		action.ActionIcon(c, lucide.Settings(), action.ActionIconProps{Title: "Borderless"})
		action.ActionIcon(c, lucide.Settings(), action.ActionIconProps{Title: "Filled", Variant: action.ActionIconFilled})
		action.ActionIcon(c, lucide.Settings(), action.ActionIconProps{Title: "Outlined", Variant: action.ActionIconOutlined})
	})
}
