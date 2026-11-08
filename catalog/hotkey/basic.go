package hotkey

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Children(func() {
		for _, platform := range []display.HotkeyPlatform{display.HotkeyApple, display.HotkeyOther} {
			ui.Row(c.Key(platform)).Gap(12).AlignItems(ui.Center).Children(func() {
				for _, keys := range []string{"mod+k", "mod+shift+p", "alt+enter", "ctrl+tab", "up", "esc"} {
					display.Hotkey(c.Key(keys), keys, display.HotkeyProps{Platform: platform})
				}
			})
		}
	})
}
