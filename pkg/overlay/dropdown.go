package overlay

import (
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/motion"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

type MenuItem struct {
	// Key is what DropdownMenu returns in the frame the item is chosen.
	Key string
	// Label is the item's text.
	Label string
	// Icon shows before the label.
	Icon *ui.SVG
	// Extra shows at the end of the row, as a shortcut.
	Extra string
	// Danger colors the item with the error color.
	Danger bool
	// Disabled dims the item and keeps it from being chosen.
	Disabled bool
	// Divider makes the item a line between groups; its other fields are ignored.
	Divider bool
}

type DropdownMenuProps struct {
	// Items are the menu's rows, in order.
	Items []MenuItem
}

var menuItemTransition = ui.ElementTransition{Duration: 150 * time.Millisecond, Ease: motion.EaseOut, Colors: true}

// DropdownMenu shows a menu of p.Items below a trigger while *open is true. trigger gives the
// trigger, a button without a look, its look and content, as action.StyleButton does. Pressing the
// trigger, or Enter, Space or Down on it, opens the menu; it returns the Key of the item chosen in
// the frame, which closes the menu, or "".
func DropdownMenu(c *ui.Context, open *bool, p DropdownMenuProps, trigger func(t ui.Element)) string {
	tk := theme.TokensOf(c)
	m := ui.MenuBase(c, open)
	trigger(m.Trigger)
	chosen := ""
	m.Popup(func(panel ui.Element) {
		panel.Margin(6, 0, 0, 0).Padding(4).Radius(tk.BorderRadius).MinWidth(220).
			Background(tk.BgElevated).Transition(motion.Popup()).AttachedOrigin()
		tk.Elevate(panel, theme.BoxShadowSecondary)
		panel.Children(func() {
			for i, it := range p.Items {
				if it.Divider {
					ui.Box(c.Key(i)).Height(1).Margin(4, 0).AlignSelf(ui.Stretch).Background(tk.Border)
					continue
				}
				if menuItem(c, tk, m, it).Clicked() && !it.Disabled {
					chosen = it.Key
				}
			}
		})
	})
	return chosen
}

func menuItem(c *ui.Context, tk *theme.Tokens, m ui.MenuParts, it MenuItem) ui.Element {
	item := m.Item().Disabled(it.Disabled).
		PaddingY(6).PaddingX(12).MinHeight(32).Radius(tk.BorderRadiusSM).AlignSelf(ui.Stretch).
		FontSize(14).Transition(menuItemTransition)
	fg, icon := tk.Text, tk.TextSecondary
	if it.Danger {
		fg, icon = tk.Error, tk.Error
	}
	switch {
	case it.Disabled:
		fg, icon = tk.TextDisabled, tk.TextDisabled
		item.Opacity(0.5)
	case item.Pressed():
		item.Background(tk.FillSecondary)
	case item.Highlighted() && it.Danger:
		item.Background(tk.ErrorBg)
	case item.Highlighted():
		item.Background(tk.FillTertiary)
	}
	item.TextColor(fg)
	item.Children(func() {
		ui.Row(c).Grow(1).Children(func() {
			if it.Icon != nil {
				ui.Icon(c, it.Icon).FontSize(14).TextColor(icon).Margin(0, 12, 0, 0)
			}
			ui.Text(c, it.Label).Grow(1).SingleLine()
			if it.Extra != "" {
				ui.Text(c, it.Extra).Font("monospace").FontSize(12).TextColor(tk.TextTertiary).Padding(0, 0, 0, 16).SingleLine()
			}
		})
	})
	return item
}
