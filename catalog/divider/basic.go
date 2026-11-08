package divider

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Column(c).Gap(16).Width(420).Children(func() {
		ui.Text(c, "Sign in with your work email.").TextColor(tk.TextSecondary)
		layout.Divider(c, layout.DividerProps{})
		ui.Text(c, "Or use a passkey on this device.").TextColor(tk.TextSecondary)
		layout.Divider(c, layout.DividerProps{Dashed: true})
		layout.Divider(c, layout.DividerProps{Label: "OR"})
	})
}
