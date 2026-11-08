package tag

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Variants(c *ui.Context) {
	variants := []struct {
		name    string
		variant display.TagVariant
	}{
		{"Filled", display.TagFilled},
		{"Outlined", display.TagOutlined},
		{"Borderless", display.TagBorderless},
		{"Solid", display.TagSolid},
	}
	for _, color := range []display.TagColor{"", display.TagBlue, display.TagError} {
		ui.Row(c).Gap(8).Children(func() {
			for _, v := range variants {
				display.Tag(c, v.name, display.TagProps{Variant: v.variant, Color: color})
			}
		})
	}
}
