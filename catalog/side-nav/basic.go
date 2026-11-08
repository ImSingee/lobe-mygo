package sidenav

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var tabs = []struct {
	key, title string
	icon       func() *ui.SVG
}{
	{"chat", "Chat", lucide.MessageSquare},
	{"discover", "Discover", lucide.Compass},
	{"files", "Files", lucide.Folder},
}

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Row(c).AlignSelf(ui.Stretch).Height(640).Radius(12).Clip().Border(1, tk.BorderSecondary)
	active := ui.Local(root, "tab", func() string { return "chat" })
	root.Children(func() {
		layout.SideNav(c, layout.SideNavProps{
			Avatar: func() { display.Avatar(c, "🤯", display.AvatarProps{Size: 40, Shape: display.AvatarCircle}) },
			TopActions: func() {
				for _, t := range tabs {
					b := action.ActionIcon(c.Key(t.key), t.icon(), action.ActionIconProps{Title: t.title, Size: theme.Large, Active: *active == t.key})
					if b.Clicked() {
						*active = t.key
					}
				}
			},
			BottomActions: func() {
				action.ActionIcon(c, lucide.CircleHelp(), action.ActionIconProps{Title: "Help"})
				action.ActionIcon(c, lucide.Settings(), action.ActionIconProps{Title: "Settings"})
			},
		})
		ui.Column(c).Grow(1).Center().Children(func() {
			ui.Text(c, "Tab: "+*active).TextColor(tk.TextSecondary)
		})
	})
}
