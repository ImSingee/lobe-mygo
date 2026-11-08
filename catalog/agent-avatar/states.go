package agentavatar

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func States(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Row(c).Gap(24).Children(func() {
		for i, name := range []string{"Idle", "Working", "Waiting", "Blocked", "Thinking", "Done"} {
			ui.Column(c.Key(i)).Gap(8).AlignItems(ui.Center).Children(func() {
				display.AgentAvatar(c, display.AgentAvatarProps{Seed: "LobeChat", Shape: display.AgentFace,
					State: display.AgentState(i), Size: 64, Animated: true})
				ui.Text(c, name).FontSize(12).TextColor(tk.TextSecondary)
			})
		}
	})
}
