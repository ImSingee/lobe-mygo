package toast

import (
	"strconv"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var archived, restored int

func Action(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Column(c).Gap(12).AlignItems(ui.Center).Children(func() {
		if action.Button(c, "Archive conversation", action.ButtonProps{}).Clicked() {
			archived++
			feedback.ShowToast(c, feedback.ToastProps{
				Title: "Conversation archived", Action: "Undo", Timeout: -1,
				OnAction: func() { restored++ },
			})
		}
		ui.Text(c, "Undone "+strconv.Itoa(restored)+" of "+strconv.Itoa(archived)).TextColor(tk.TextSecondary)
	})
	feedback.ToastHost(c, feedback.ToastHostProps{})
}
