package button

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func Sizes(c *ui.Context) {
	for _, size := range []theme.Size{theme.Small, theme.Middle, theme.Large} {
		ui.Row(c).Gap(8).Children(func() {
			action.Button(c, "Primary", action.ButtonProps{Variant: action.Primary, Size: size})
			action.Button(c, "Default", action.ButtonProps{Size: size})
			action.Button(c, "Filled", action.ButtonProps{Variant: action.Filled, Size: size})
		})
	}
}
