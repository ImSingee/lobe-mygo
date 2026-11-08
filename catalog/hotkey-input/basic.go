package hotkeyinput

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
)

var (
	search  = "mod+k"
	newChat = ""
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Width(320).Children(func() {
		row := func(label string, fn func()) {
			ui.Column(c.Key(label)).Gap(6).Children(func() {
				ui.Text(c, label).FontSize(12)
				fn()
			})
		}
		row("Search", func() {
			field.HotkeyInput(c, &search, field.HotkeyInputProps{ResetValue: "mod+k", AllowClear: true, Conflicts: []string{"mod+s"}}).Label("Search")
		})
		row("New chat", func() {
			field.HotkeyInput(c, &newChat, field.HotkeyInputProps{Variant: field.InputFilled, AllowClear: true}).Label("New chat")
		})
	})
}
