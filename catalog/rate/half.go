package rate

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
)

var half = 2.5

func Half(c *ui.Context) {
	ui.Column(c).Gap(12).Children(func() {
		choice.Rate(c, &half, choice.RateProps{AllowHalf: true, Size: 28}).Label("Half stars")
		ui.Text(c, fmt.Sprintf("%g out of 5", half)).FontSize(12)
	})
}
