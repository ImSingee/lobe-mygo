package contextmenu

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/overlay"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var lastPick = "nothing yet"

var fileItems = []overlay.MenuItem{
	{Key: "open", Label: "Open", Icon: lucide.FolderOpen()},
	{Key: "rename", Label: "Rename", Icon: lucide.Pencil(), Hotkey: "f2"},
	{Key: "duplicate", Label: "Duplicate", Icon: lucide.Copy(), Hotkey: "mod+d"},
	{Divider: true},
	{Key: "delete", Label: "Delete", Icon: lucide.Trash(), Danger: true},
}

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Column(c).Gap(12).AlignItems(ui.Center).Children(func() {
		area := ui.Column(c).Size(360, 160).Radius(tk.BorderRadiusLG).Border(1, tk.Border).BorderStyle(ui.BorderDashed).
			Justify(ui.Center).AlignItems(ui.Center).Focusable().Label("Right-click area")
		area.Children(func() {
			ui.Text(c, "Right-click here").TextColor(tk.TextSecondary)
		})
		if key := overlay.ContextMenu(c, area, overlay.DropdownMenuProps{Items: fileItems}); key != "" {
			lastPick = key
		}
		ui.Text(c, "Last chosen: "+lastPick).TextColor(tk.TextSecondary)
	})
}
