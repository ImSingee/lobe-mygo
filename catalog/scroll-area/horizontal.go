package scrollarea

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/layout"
)

var topics = []string{"Design", "Engineering", "Research", "Marketing", "Sales", "Support", "Finance", "Legal", "Operations", "People", "Data", "Security"}

func Horizontal(c *ui.Context) {
	layout.ScrollArea(c, layout.ScrollAreaProps{Direction: layout.Horizontal, Fade: true}, func() {
		ui.Row(c).Gap(8).Padding(4, 4, 16, 4).Children(func() {
			for _, topic := range topics {
				display.Tag(c.Key(topic), topic, display.TagProps{})
			}
		})
	}).Width(360).Height(48)
}
