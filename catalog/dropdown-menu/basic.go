package dropdownmenu

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/overlay"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var (
	menuOpen bool
	lastPick = "nothing yet"
)

var fileItems = []overlay.MenuItem{
	{Key: "rename", Label: "Rename", Icon: lucide.Pencil(), Hotkey: "f2"},
	{Key: "duplicate", Label: "Duplicate", Icon: lucide.Copy(), Hotkey: "mod+d"},
	{Key: "share", Label: "Share", Icon: lucide.Share2(), Disabled: true},
	{Divider: true},
	{Key: "archive", Label: "Archive", Icon: lucide.Archive()},
	{Key: "delete", Label: "Delete", Icon: lucide.Trash(), Danger: true},
}

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Column(c).Gap(12).AlignItems(ui.Center).Children(func() {
		key := overlay.DropdownMenu(c, &menuOpen, overlay.DropdownMenuProps{Items: fileItems}, func(t ui.Element) {
			action.StyleButton(c, t, "Actions", action.ButtonProps{}).Children(func() {
				ui.Icon(c, lucide.ChevronDown()).FontSize(14).TextColor(tk.TextTertiary)
			})
		})
		if key != "" {
			lastPick = key
		}
		ui.Text(c, "Last chosen: "+lastPick).TextColor(tk.TextSecondary)
	})
}
