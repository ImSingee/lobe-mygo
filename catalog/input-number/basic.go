package inputnumber

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var (
	quantity = 3.0
	price    = 19.9
	small    = 1.0
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(12).Width(240).Children(func() {
		field.InputNumber(c, &quantity, field.InputNumberProps{Min: 1, Max: 10}).Label("Quantity")
		field.InputNumber(c, &price, field.InputNumberProps{Step: 0.1, Precision: 2, Variant: field.InputFilled}).Label("Price")
		field.InputNumber(c, &small, field.InputNumberProps{Size: theme.Small}).Label("Small")
		ui.Text(c, fmt.Sprintf("Quantity %g · Price %.2f", quantity, price)).FontSize(12)
	})
}
