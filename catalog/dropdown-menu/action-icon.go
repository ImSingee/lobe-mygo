package dropdownmenu

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/overlay"
)

var moreOpen bool

var moreItems = []overlay.MenuItem{
	{Key: "copy-link", Label: "Copy link", Icon: lucide.Link()},
	{Key: "search", Label: "Find in page", Icon: lucide.Search(), Extra: "⌘F"},
}

func ActionIconTrigger(c *ui.Context) {
	overlay.DropdownMenu(c, &moreOpen, overlay.DropdownMenuProps{Items: moreItems}, func(t ui.Element) {
		action.StyleActionIcon(c, t, lucide.Ellipsis(), action.ActionIconProps{Active: moreOpen}).Label("More")
	})
}
