package tabs

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/nav"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var variantTabs = [5]int{}

func Variants(c *ui.Context) {
	plain := []nav.TabItem{{Label: "Overview"}, {Label: "Activity"}, {Label: "Members"}}
	ui.Column(c).Gap(20).Width(420).Children(func() {
		nav.Tabs(c.Key("small"), &variantTabs[0], nav.TabsProps{Items: plain, Size: theme.Small}, nil)
		nav.Tabs(c.Key("large"), &variantTabs[1], nav.TabsProps{Items: plain, Size: theme.Large}, nil)
		nav.Tabs(c.Key("square"), &variantTabs[2], nav.TabsProps{Items: plain, Variant: nav.TabsSquare}, nil)
		nav.Tabs(c.Key("point"), &variantTabs[3], nav.TabsProps{Items: plain, Variant: nav.TabsPoint}, nil)
	})
}
