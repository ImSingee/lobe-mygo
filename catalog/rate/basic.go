package rate

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

var (
	stars  = 3.0
	hearts = 2.0
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Children(func() {
		choice.Rate(c, &stars, choice.RateProps{}).Label("Rating")
		choice.Rate(c, &hearts, choice.RateProps{Character: lucide.Heart(), Color: ui.Hex("#f5222d"), Size: 24}).Label("Love")
		fixed := 4.0
		choice.Rate(c, &fixed, choice.RateProps{Disabled: true}).Label("Read only")
		ui.Text(c, fmt.Sprintf("%g stars · %g hearts", stars, hearts)).FontSize(12)
	})
}
