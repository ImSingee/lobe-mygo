package segmented

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var picks = make([]int, 4)

func Variants(c *ui.Context) {
	options := []string{"List", "Board", "Calendar"}
	ui.Column(c).Gap(12).Width(360).AlignItems(ui.Center).Children(func() {
		action.Segmented(c, &picks[0], options, action.SegmentedProps{Variant: action.SegmentedOutlined})
		action.Segmented(c, &picks[1], options, action.SegmentedProps{Size: theme.Small})
		action.Segmented(c, &picks[2], options, action.SegmentedProps{Size: theme.Large})
		action.Segmented(c, &picks[3], options, action.SegmentedProps{Block: true})
	})
}
