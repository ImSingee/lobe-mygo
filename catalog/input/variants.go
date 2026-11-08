package input

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
)

var variantValues = make([]string, 3)

func Variants(c *ui.Context) {
	ui.Column(c).Gap(12).Width(320).Children(func() {
		field.Input(c, &variantValues[0], field.InputProps{Placeholder: "Outlined"})
		field.Input(c, &variantValues[1], field.InputProps{Placeholder: "Filled", Variant: field.InputFilled})
		field.Input(c, &variantValues[2], field.InputProps{Placeholder: "Borderless", Variant: field.InputBorderless})
	})
}
