package burger

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/nav"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var items = []display.ListItem{
	{Key: "chat", Label: "Chat", Icon: lucide.MessageSquare()},
	{Key: "discover", Label: "Discover", Icon: lucide.Compass()},
	{Key: "files", Label: "Files", Icon: lucide.Folder()},
	{Divider: true},
	{Key: "settings", Label: "Settings", Icon: lucide.Settings()},
}

type state struct {
	open, fullscreen bool
	active           string
}

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Column(c).Absolute().Left(0).Top(0).Right(0).Bottom(0)
	st := ui.Local(root, "burger", func() state { return state{active: "chat"} })
	root.Children(func() {
		ui.Row(c).Height(64).PaddingX(16).AlignItems(ui.Center).Gap(8).Shrink(0).BorderWidth(0, 0, 1, 0).BorderColor(tk.BorderSecondary).Children(func() {
			ui.Text(c, "LobeHub").FontSize(16).FontWeight(700).Grow(1)
			nav.Burger(c, &st.open, nav.BurgerProps{Items: items, Active: st.active, Fullscreen: st.fullscreen,
				OnSelect: func(k string) { st.active = k },
				Footer: func() {
					action.Button(c, "Sign in", action.ButtonProps{Variant: action.Primary}).Grow(1)
				}})
		})
		ui.Column(c).Padding(24).Gap(12).AlignItems(ui.Start).Children(func() {
			ui.Text(c, "Page: "+st.active).TextColor(tk.TextSecondary)
			label := "Fullscreen menu"
			if st.fullscreen {
				label = "Drawer under the header"
			}
			if action.Button(c, label, action.ButtonProps{}).Clicked() {
				st.fullscreen = !st.fullscreen
			}
		})
	})
}
