package tag

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func Sizes(c *ui.Context) {
	for _, round := range []bool{false, true} {
		ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
			display.Tag(c, "Small", display.TagProps{Size: theme.Small, Round: round, Color: display.TagGeekblue})
			display.Tag(c, "Middle", display.TagProps{Size: theme.Middle, Round: round, Color: display.TagGeekblue})
			display.Tag(c, "Large", display.TagProps{Size: theme.Large, Round: round, Color: display.TagGeekblue})
		})
	}
}
