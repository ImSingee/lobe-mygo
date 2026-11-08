package home

import (
	"slices"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

type showcase struct {
	width     float32
	expanded  bool
	active    int
	convs     []conversation
	draft     string
	streaming int
	start     time.Time
	sent      int
}

func newShowcase() showcase {
	convs := make([]conversation, len(conversations))
	for i, cv := range conversations {
		cv.messages = slices.Clone(cv.messages)
		convs[i] = cv
	}
	return showcase{width: 272, expanded: true, convs: convs, streaming: len(convs[0].messages) - 1}
}

// Showcase is the home page's live app: a chat client built from Lobe MyGo components.
func Showcase(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Row(c).Absolute().Left(0).Top(0).Right(0).Bottom(0).AlignItems(ui.Stretch).Background(tk.BgContainer)
	st := ui.Local(root, "showcase", newShowcase)
	if st.start.IsZero() {
		st.start = c.Now()
	}
	root.Children(func() {
		layout.DraggableSideNav(c, &st.width, &st.expanded, layout.DraggableSideNavProps{
			MinWidth: 54,
			MaxWidth: 320,
			Header:   func(nav layout.SideNavState) { sideHeader(c, tk, nav) },
			Footer: func(layout.SideNavState) {
				ui.Row(c).Padding(8, 9).Children(func() {
					action.ActionIcon(c, lucide.Settings(), action.ActionIconProps{Title: "Settings"})
				})
			},
		}, func(nav layout.SideNavState) { agentList(c, st, nav) })
		chatPane(c, tk, st)
	})
}

func sideHeader(c *ui.Context, tk *theme.Tokens, nav layout.SideNavState) {
	ui.Row(c).Height(56).Padding(0, 9).Gap(10).AlignItems(ui.Center).Children(func() {
		display.Avatar(c, "🤯", display.AvatarProps{Size: 36, Background: tk.FillTertiary})
		if nav.Expanded {
			ui.Row(c).Grow(1).Basis(0).Gap(4).AlignItems(ui.Center).Opacity(1 - nav.Collapse).Children(func() {
				ui.Text(c, "LobeHub").FontWeight(600).SingleLine().Grow(1).Basis(0)
				action.ActionIcon(c, lucide.SquarePen(), action.ActionIconProps{Title: "New chat", Size: theme.Small})
			})
		}
	})
}

func agentAvatar(c *ui.Context, cv conversation, size float32, busy bool) ui.Element {
	state := display.AgentIdle
	if busy {
		state = display.AgentWorking
	}
	return display.AgentAvatar(c, display.AgentAvatarProps{Seed: cv.name, Shape: cv.shape, Color: cv.color,
		Size: size, State: state, Animated: true})
}

func agentList(c *ui.Context, st *showcase, nav layout.SideNavState) {
	items := make([]display.ListItem, len(st.convs))
	for i, cv := range st.convs {
		busy := i == st.active && st.streaming >= 0
		items[i] = display.ListItem{Key: cv.name, Label: cv.name, Description: cv.topic,
			Avatar: func() { agentAvatar(c.Key(cv.name), cv, 36, busy) }}
	}
	ui.Column(c).Padding(0, 0, 8, 0).Children(func() {
		k := display.List(c, display.ListProps{Items: items, Active: st.convs[st.active].name,
			Collapse: nav.Collapse, RailWidth: nav.RailWidth})
		for i, cv := range st.convs {
			if k == cv.name && i != st.active {
				st.active, st.streaming = i, -1
			}
		}
	})
}
