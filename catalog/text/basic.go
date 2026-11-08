package text

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(8).Children(func() {
		display.Text(c, "Lobe UI (default)", display.TextProps{})
		display.Text(c, "Lobe UI (secondary)", display.TextProps{Type: display.TextSecondary})
		display.Text(c, "Lobe UI (success)", display.TextProps{Type: display.TextSuccess})
		display.Text(c, "Lobe UI (warning)", display.TextProps{Type: display.TextWarning})
		display.Text(c, "Lobe UI (danger)", display.TextProps{Type: display.TextDanger})
		display.Text(c, "Lobe UI (disabled)", display.TextProps{Disabled: true})
		ui.Row(c).Gap(16).Children(func() {
			display.Text(c, "strong", display.TextProps{Strong: true})
			display.Text(c, "italic", display.TextProps{Italic: true})
			display.Text(c, "underline", display.TextProps{Underline: true})
			display.Text(c, "delete", display.TextProps{Delete: true})
			display.Text(c, "code", display.TextProps{Code: true})
			display.Text(c, "mark", display.TextProps{Mark: true})
		})
		display.Text(c, "Thinking…", display.TextProps{Shiny: true})
	})
}
