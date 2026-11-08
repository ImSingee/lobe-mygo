package icon

import (
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
	"github.com/lobehub/lobe-mygo/lucide"
)

func Spin(c *ui.Context) {
	ui.Row(c).Gap(16).Children(func() {
		lobe.Icon(c, lucide.LoaderCircle(), lobe.IconProps{Size: lobe.IconLarge, Spin: true})
		lobe.Icon(c, lucide.RefreshCw(), lobe.IconProps{Size: lobe.IconLarge, Spin: true})
	})
}
