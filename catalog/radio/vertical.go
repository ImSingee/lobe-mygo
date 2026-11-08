package radio

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
)

var density = 1

func Vertical(c *ui.Context) {
	choice.RadioGroup(c, choice.RadioGroupProps{Vertical: true}, func() {
		for i, label := range []string{"Compact", "Comfortable", "Spacious"} {
			choice.Radio(c.Key(i), &density, i, label, choice.RadioProps{Size: 20})
		}
	}).Label("Density")
}
