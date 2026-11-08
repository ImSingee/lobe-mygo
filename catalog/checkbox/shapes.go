package checkbox

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
)

var shapes = []bool{true, true, false, true}

func Shapes(c *ui.Context) {
	ui.Row(c).Gap(20).Children(func() {
		choice.Checkbox(c, &shapes[0], "Square", choice.CheckboxProps{})
		choice.Checkbox(c, &shapes[1], "Circle", choice.CheckboxProps{Shape: choice.CheckboxCircle})
		choice.Checkbox(c, &shapes[2], "Large", choice.CheckboxProps{Size: 20})
		choice.Checkbox(c, &shapes[3], "Disabled", choice.CheckboxProps{}).Disabled(true)
	})
}
