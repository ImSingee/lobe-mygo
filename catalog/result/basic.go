package result

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
)

var status = 1

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Width(420).AlignItems(ui.Center).Children(func() {
		action.Segmented(c, &status, []string{"Info", "Success", "Warning", "Error"}, action.SegmentedProps{})
		texts := [][2]string{
			{"Your workspace is ready", "Invite your team to start building agents together."},
			{"Payment received", "A receipt for $20.00 is on its way to your inbox."},
			{"Your plan renews in 3 days", "Update your payment method to keep your agents running."},
			{"The deploy failed", "The build ran out of memory. Check the logs and try again."},
		}
		feedback.Result(c, feedback.ResultProps{
			Status: feedback.ResultStatus(status), Title: texts[status][0], SubTitle: texts[status][1],
			Extra: func() {
				action.Button(c, "Continue", action.ButtonProps{Variant: action.Primary})
				action.Button(c, "Back", action.ButtonProps{})
			},
		})
	})
}
