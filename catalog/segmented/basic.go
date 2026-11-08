package segmented

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
)

var period int

func Basic(c *ui.Context) {
	action.Segmented(c, &period, []string{"Daily", "Weekly", "Monthly", "Yearly"}, action.SegmentedProps{})
}
