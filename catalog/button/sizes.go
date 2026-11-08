package button

import (
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
)

func Sizes(c *ui.Context) {
	for _, size := range []lobe.Size{lobe.Small, lobe.Middle, lobe.Large} {
		ui.Row(c).Gap(8).Children(func() {
			lobe.Button(c, "Primary", lobe.ButtonProps{Variant: lobe.Primary, Size: size})
			lobe.Button(c, "Default", lobe.ButtonProps{Size: size})
			lobe.Button(c, "Filled", lobe.ButtonProps{Variant: lobe.Filled, Size: size})
		})
	}
}
