package accordion

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

var variantOpen = []map[string]bool{{"general": true}, {"general": true}, {"general": true}}

var settings = []display.AccordionItem{
	{Key: "general", Title: "General"},
	{Key: "models", Title: "Models"},
	{Key: "billing", Title: "Billing", Disabled: true},
}

func Variants(c *ui.Context) {
	ui.Column(c).Gap(24).Width(480).Children(func() {
		for i, v := range []struct {
			name      string
			variant   display.AccordionVariant
			indicator display.AccordionIndicator
			gap       float32
		}{
			{"Borderless", display.AccordionBorderless, display.AccordionIndicatorStart, 0},
			{"Filled", display.AccordionFilled, display.AccordionIndicatorEnd, 4},
			{"Outlined", display.AccordionOutlined, display.AccordionIndicatorStart, 0},
		} {
			ui.Column(c.Key(v.name)).Gap(8).Children(func() {
				ui.Text(c, v.name).FontSize(12)
				display.Accordion(c, variantOpen[i], display.AccordionProps{Items: settings, Variant: v.variant, Indicator: v.indicator, Gap: v.gap}, func(key string) {
					ui.Text(c, "Settings for "+key+" go here.")
				})
			})
		}
	})
}
