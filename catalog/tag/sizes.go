package tag

import (
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
)

func Sizes(c *ui.Context) {
	for _, round := range []bool{false, true} {
		ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
			lobe.Tag(c, "Small", lobe.TagProps{Size: lobe.Small, Round: round, Color: lobe.TagGeekblue})
			lobe.Tag(c, "Middle", lobe.TagProps{Size: lobe.Middle, Round: round, Color: lobe.TagGeekblue})
			lobe.Tag(c, "Large", lobe.TagProps{Size: lobe.Large, Round: round, Color: lobe.TagGeekblue})
		})
	}
}
