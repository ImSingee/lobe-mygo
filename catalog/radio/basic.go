package radio

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
)

var plan = "pro"

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Children(func() {
		choice.RadioGroup(c, choice.RadioGroupProps{}, func() {
			choice.Radio(c, &plan, "free", "Free", choice.RadioProps{})
			choice.Radio(c, &plan, "pro", "Pro", choice.RadioProps{})
			choice.Radio(c, &plan, "team", "Team", choice.RadioProps{})
			choice.Radio(c, &plan, "enterprise", "Enterprise", choice.RadioProps{Disabled: true})
		}).Label("Plan")
		ui.Text(c, "Plan: "+plan).FontSize(12)
	})
}
