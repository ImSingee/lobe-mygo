package avatar

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Basic(c *ui.Context) {
	ui.Row(c).Gap(12).Children(func() {
		for _, name := range []string{"Ada Lovelace", "Grace Hopper", "Alan Turing", "Lobe", "Claude"} {
			display.Avatar(c.Key(name), name, display.AvatarProps{})
		}
	})
}
