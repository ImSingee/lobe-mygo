package togglegroup

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var (
	views = [][]string{{"list"}, {"list"}, {"list"}, {"list"}}
	items = []choice.ToggleItem{{Value: "list", Label: "List"}, {Value: "board", Label: "Board"}, {Value: "calendar", Label: "Calendar"}}
)

func Variants(c *ui.Context) {
	ui.Column(c).Gap(16).Children(func() {
		for i, v := range []struct {
			variant choice.ToggleGroupVariant
			size    theme.Size
		}{{choice.ToggleOutlined, theme.Middle}, {choice.ToggleBorderless, theme.Middle}, {choice.ToggleSegmented, theme.Middle}, {choice.ToggleOutlined, theme.Small}} {
			choice.ToggleGroup(c.Key(i), &views[i], choice.ToggleGroupProps{Items: items, Variant: v.variant, Size: v.size})
		}
	})
}
