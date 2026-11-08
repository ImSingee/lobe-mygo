package steps

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/nav"
)

func Vertical(c *ui.Context) {
	ui.Row(c).Gap(48).AlignItems(ui.Start).Children(func() {
		ui.Column(c).Width(220).Children(func() {
			nav.Steps(c.Key("default"), nav.StepsProps{Vertical: true, Current: 2, Items: []nav.StepItem{
				{Title: "Upload", Description: "report.pdf, 2.4 MB"},
				{Title: "Parse", Description: "Split into 38 chunks"},
				{Title: "Embed", Description: "Failed to reach the model", Status: nav.StepError},
				{Title: "Index"},
			}})
		})
		ui.Column(c).Width(220).Children(func() {
			nav.Steps(c.Key("dot"), nav.StepsProps{Vertical: true, Dot: true, Current: 1, Items: []nav.StepItem{
				{Title: "Ordered", Description: "Nov 6"},
				{Title: "Shipped", Description: "Nov 7"},
				{Title: "Delivered"},
			}})
		})
	})
}
