package draggablesidenav

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var agents = []struct{ avatar, name, subtitle string }{
	{"👨‍💼", "LiJian", ""},
	{"👩‍💼", "Rene Wang", "Playing Cursor"},
	{"🤖", "Carl-bot", "Playing /games | carl.gg"},
	{"👔", "Arvin Xu", ""},
	{"🎨", "cy948", ""},
	{"💻", "Wuxh", "Coding"},
	{"🎬", "Brandon Studio", ""},
	{"🌙", "Tsuki", ""},
}

type state struct {
	width    float32
	expanded bool
	active   string
}

func header(c *ui.Context, st *state, nav layout.SideNavState) {
	ui.Column(c).Children(func() {
		ui.Row(c).Padding(4).Margin(4).Gap(8).AlignItems(ui.Center).Children(func() {
			display.Avatar(c, "🐶", display.AvatarProps{Size: 36, Background: theme.TokensOf(c).Blue})
			if nav.Expanded {
				ui.Row(c).Grow(1).Basis(0).Gap(8).AlignItems(ui.Center).Opacity(1 - nav.Collapse).Children(func() {
					ui.Text(c, "Canis workspace").SingleLine().Grow(1).Basis(0)
					action.ActionIcon(c, lucide.ChevronDown(), action.ActionIconProps{Size: theme.Small})
				})
			}
		})
		items := []display.ListItem{
			{Key: "home", Label: "Home", Icon: lucide.House()},
			{Key: "integrations", Label: "Integrations", Icon: lucide.SquareDashedBottom()},
			{Key: "community", Label: "Community", Icon: lucide.Users()},
		}
		if k := display.List(c, rail(items, st, nav)); k != "" {
			st.active = k
		}
	})
}

func rail(items []display.ListItem, st *state, nav layout.SideNavState) display.ListProps {
	return display.ListProps{Items: items, Active: st.active, Collapse: nav.Collapse, RailWidth: nav.RailWidth}
}

func body(c *ui.Context, st *state, nav layout.SideNavState) {
	items := []display.ListItem{
		{Key: "repo-1", Label: "Repo 1", Icon: lucide.Folder()},
		{Key: "repo-2", Label: "Repo 2", Icon: lucide.Folder()},
		{Divider: true},
	}
	for _, a := range agents {
		items = append(items, display.ListItem{Key: a.name, Label: a.name, Description: a.subtitle,
			Avatar: func() { display.Avatar(c, a.avatar, display.AvatarProps{Size: 36}) }})
	}
	if k := display.List(c, rail(items, st, nav)); k != "" {
		st.active = k
	}
}

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Row(c).Absolute().Left(0).Top(0).Right(0).Bottom(0).AlignItems(ui.Stretch)
	st := ui.Local(root, "nav", func() state { return state{width: 300, expanded: true, active: "home"} })
	root.Children(func() {
		layout.DraggableSideNav(c, &st.width, &st.expanded, layout.DraggableSideNavProps{
			MaxWidth: 320,
			MinWidth: 54,
			Header:   func(nav layout.SideNavState) { header(c, st, nav) },
			Footer:   func(layout.SideNavState) { ui.Row(c).Padding(8) },
		}, func(nav layout.SideNavState) { body(c, st, nav) })
		ui.Column(c).Grow(1).Basis(0).Padding(24).Gap(16).AlignItems(ui.Start).Children(func() {
			ui.Column(c).Gap(8).Children(func() {
				ui.Text(c, "DraggableSideNav Demo").FontSize(20).FontWeight(700)
				ui.Text(c, "A workspace-style side panel with draggable resize").TextColor(tk.TextSecondary)
			})
			ui.Text(c, "Active: "+st.active).FontWeight(500)
			label := "Collapse sidebar"
			if !st.expanded {
				label = "Expand sidebar"
			}
			if action.Button(c, label, action.ButtonProps{Variant: action.Primary}).Clicked() {
				st.expanded = !st.expanded
			}
			ui.Text(c, "Hover the edge to reveal the toggle, drag it to resize, or drag it under a third of the way to collapse.").
				FontSize(12).TextColor(tk.TextTertiary)
		})
	})
}
