package actionicongroup

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/overlay"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var lastAction = "nothing yet"

var (
	items = []action.ActionGroupItem{
		{Key: "copy", Icon: lucide.Copy(), Label: "Copy"},
		{Key: "edit", Icon: lucide.Pencil(), Label: "Edit"},
		{Key: "regenerate", Icon: lucide.RotateCw(), Label: "Regenerate"},
		{Key: "delete", Icon: lucide.Trash(), Label: "Delete", Danger: true},
	}
	menu = []overlay.MenuItem{
		{Key: "share", Label: "Share", Icon: lucide.Share2()},
		{Key: "export", Label: "Export as Markdown", Icon: lucide.Download(), Hotkey: "mod+e"},
	}
)

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	group := func(key string, p action.ActionIconGroupProps) {
		p.Items, p.Menu, p.Size = items, menu, theme.Small
		if k := action.ActionIconGroup(c.Key(key), p); k != "" {
			lastAction = k
		}
	}
	ui.Column(c).Gap(16).AlignItems(ui.Center).Children(func() {
		ui.Row(c).Gap(24).AlignItems(ui.Start).Children(func() {
			group("filled", action.ActionIconGroupProps{})
			group("outlined", action.ActionIconGroupProps{Variant: action.GroupOutlined, Shadow: true})
			group("borderless", action.ActionIconGroupProps{Variant: action.GroupBorderless})
			group("vertical", action.ActionIconGroupProps{Vertical: true})
		})
		ui.Text(c, "Last action: "+lastAction).TextColor(tk.TextSecondary)
	})
}
