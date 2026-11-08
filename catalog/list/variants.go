package list

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

var settingsActive = "general"

func Variants(c *ui.Context) {
	items := []display.ListItem{
		{Key: "general", Label: "General", Icon: lucide.Settings()},
		{Key: "models", Label: "Models", Icon: lucide.Brain(), Extra: "12"},
		{Key: "storage", Label: "Storage", Icon: lucide.HardDrive(), Disabled: true},
		{Divider: true},
		{Key: "signout", Label: "Sign out", Icon: lucide.LogOut(), Danger: true},
	}
	ui.Row(c).Gap(24).AlignItems(ui.Start).Children(func() {
		ui.Column(c).Width(220).Children(func() {
			if k := display.List(c.Key("filled"), display.ListProps{Active: settingsActive, Items: items, Variant: display.ListFilled}); k != "" {
				settingsActive = k
			}
		})
		ui.Column(c).Width(220).Children(func() {
			if k := display.List(c.Key("outlined"), display.ListProps{Active: settingsActive, Items: items, Variant: display.ListOutlined, Compact: true}); k != "" {
				settingsActive = k
			}
		})
	})
}
