package hotkey

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Variants(c *ui.Context) {
	ui.Column(c).Gap(16).Children(func() {
		for _, v := range []display.HotkeyVariant{display.HotkeyFilled, display.HotkeyOutlined, display.HotkeyBorderless} {
			ui.Row(c.Key(v)).Gap(12).AlignItems(ui.Center).Children(func() {
				display.Hotkey(c, "mod+shift+k", display.HotkeyProps{Variant: v})
				display.Hotkey(c, "mod+shift+k", display.HotkeyProps{Variant: v, Compact: true})
			})
		}
	})
}
