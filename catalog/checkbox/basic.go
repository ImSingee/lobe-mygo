package checkbox

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
)

var agreed bool

func Basic(c *ui.Context) {
	choice.Checkbox(c, &agreed, "I agree to the terms", choice.CheckboxProps{})
}
