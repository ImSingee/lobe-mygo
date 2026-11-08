package agentavatar

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

var agents = []string{"Research", "Writer", "Coder", "Designer", "Translator", "Planner", "Reviewer", "Tutor", "Analyst", "Scout", "Editor", "Critic"}

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Children(func() {
		ui.Row(c).Gap(12).Wrap().Width(520).Children(func() {
			for s := display.AgentFace; s <= display.AgentCat; s++ {
				display.AgentAvatar(c.Key(int(s)), display.AgentAvatarProps{Shape: s, Seed: agents[int(s)-1], Size: 56, Animated: true})
			}
		})
		ui.Row(c).Gap(12).Children(func() {
			for i, name := range agents[:7] {
				display.AgentAvatar(c.Key(name), display.AgentAvatarProps{Name: name, Shape: display.AgentFace,
					Color: display.AgentColor(i + 1), Size: 40, NoTilt: true})
			}
		})
	})
}
