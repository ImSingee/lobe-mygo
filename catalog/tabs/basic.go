package tabs

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/nav"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var (
	basicTab int
	items    = []nav.TabItem{
		{Label: "Chat", Icon: lucide.MessageSquare()},
		{Label: "Files", Icon: lucide.Folder()},
		{Label: "Settings", Icon: lucide.Settings()},
		{Label: "Billing", Icon: lucide.CreditCard(), Disabled: true},
	}
	panels = []string{
		"Talk to your agents and keep every conversation in one place.",
		"Upload documents for the agents to read and cite.",
		"Pick the model, the language and how replies are shown.",
		"",
	}
)

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Column(c).Width(420).Children(func() {
		nav.Tabs(c, &basicTab, nav.TabsProps{Items: items}, func(i int) {
			ui.Text(c, panels[i]).TextColor(tk.TextSecondary)
		})
	})
}
