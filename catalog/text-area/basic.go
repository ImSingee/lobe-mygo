package textarea

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
)

var notes, feedback string

func Basic(c *ui.Context) {
	ui.Column(c).Gap(12).Width(360).Children(func() {
		field.TextArea(c, &notes, field.TextAreaProps{Placeholder: "Grows from 2 to 6 lines", MaxRows: 6})
		field.TextArea(c, &feedback, field.TextAreaProps{Placeholder: "Filled", Variant: field.InputFilled})
	})
}
