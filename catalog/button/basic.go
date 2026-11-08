package button

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
)

var saved int

func Basic(c *ui.Context) {
	ui.Row(c).Gap(8).Children(func() {
		if lobe.Button(c, "Reset", lobe.ButtonProps{}).Clicked() {
			saved = 0
		}
		if lobe.Button(c, "Save", lobe.ButtonProps{Variant: lobe.Primary}).Clicked() {
			saved++
		}
	})
	ui.Text(c, fmt.Sprintf("Saved %d times", saved)).TextColor(lobe.TokensOf(c).TextSecondary)
}
