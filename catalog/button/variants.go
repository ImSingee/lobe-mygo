package button

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
)

func Variants(c *ui.Context) {
	ui.Row(c).Gap(8).Wrap().Children(func() {
		action.Button(c, "Primary", action.ButtonProps{Variant: action.Primary})
		action.Button(c, "Default", action.ButtonProps{Variant: action.Default})
		action.Button(c, "Dashed", action.ButtonProps{Variant: action.Dashed})
		action.Button(c, "Filled", action.ButtonProps{Variant: action.Filled})
		action.Button(c, "Text", action.ButtonProps{Variant: action.Text})
		action.Button(c, "Link", action.ButtonProps{Variant: action.Link})
	})
}
