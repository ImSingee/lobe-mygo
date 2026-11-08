package tag

import (
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
)

func Variants(c *ui.Context) {
	variants := []struct {
		name    string
		variant lobe.TagVariant
	}{
		{"Filled", lobe.TagFilled},
		{"Outlined", lobe.TagOutlined},
		{"Borderless", lobe.TagBorderless},
		{"Solid", lobe.TagSolid},
	}
	for _, color := range []lobe.TagColor{"", lobe.TagBlue, lobe.TagError} {
		ui.Row(c).Gap(8).Children(func() {
			for _, v := range variants {
				lobe.Tag(c, v.name, lobe.TagProps{Variant: v.variant, Color: color})
			}
		})
	}
}
