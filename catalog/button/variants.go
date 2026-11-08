package button

import (
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
)

func Variants(c *ui.Context) {
	ui.Row(c).Gap(8).Wrap().Children(func() {
		lobe.Button(c, "Primary", lobe.ButtonProps{Variant: lobe.Primary})
		lobe.Button(c, "Default", lobe.ButtonProps{Variant: lobe.Default})
		lobe.Button(c, "Dashed", lobe.ButtonProps{Variant: lobe.Dashed})
		lobe.Button(c, "Filled", lobe.ButtonProps{Variant: lobe.Filled})
		lobe.Button(c, "Text", lobe.ButtonProps{Variant: lobe.Text})
		lobe.Button(c, "Link", lobe.ButtonProps{Variant: lobe.Link})
	})
}
