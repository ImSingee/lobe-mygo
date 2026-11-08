package guidecard

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/layout"
)

func Basic(c *ui.Context) {
	ui.Row(c).Gap(16).AlignItems(ui.Start).Children(func() {
		for i, v := range []layout.BlockVariant{layout.BlockFilled, layout.BlockOutlined, layout.BlockBorderless} {
			layout.GuideCard(c.Key(i), layout.GuideCardProps{Variant: v, Shadow: v == layout.BlockOutlined,
				Title: "Meet your agents", Desc: "Agents remember your context and pick up where you left off."}).Width(220)
		}
	})
}
