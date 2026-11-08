package modal

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/overlay"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var publishOpen bool

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	if action.Button(c, "Open modal", action.ButtonProps{Variant: action.Primary}).Clicked() {
		publishOpen = true
	}
	ok := overlay.Modal(c, &publishOpen, overlay.ModalProps{Title: "Publish changes"}, func() {
		ui.Text(c, "Your edits go live for everyone with access to this page. You can roll back from the history panel.").
			TextColor(tk.TextSecondary).LineHeight(1.6)
	})
	if ok {
		publishOpen = false
	}
}
