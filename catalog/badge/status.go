package badge

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
)

func Status(c *ui.Context) {
	ui.Column(c).Gap(10).Children(func() {
		for i, s := range []struct {
			status feedback.BadgeStatus
			text   string
		}{{feedback.BadgeSuccess, "Online"}, {feedback.BadgeProcessing, "Syncing"}, {feedback.BadgeWarning, "Degraded"}, {feedback.BadgeError, "Offline"}, {feedback.BadgeDefault, "Idle"}} {
			feedback.Badge(c.Key(i), feedback.BadgeProps{Status: s.status, Text: s.text}, nil)
		}
	})
}
