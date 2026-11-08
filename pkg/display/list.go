package display

import (
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/motion"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

type ListVariant int

const (
	ListBorderless ListVariant = iota
	ListFilled
	ListOutlined
)

type ListItem struct {
	// Key is what List returns in the frame the item is clicked, and what Active names.
	Key string
	// Label is the item's text, on one line; Description shows under it in small type.
	Label, Description string
	// Icon shows before the label; Avatar, if set, builds what shows there instead.
	Icon   *ui.SVG
	Avatar func()
	// Extra shows small text at the end of the row.
	Extra string
	// Actions builds buttons that show at the end of the row while the pointer is on it or the
	// focus is in it.
	Actions func()
	// ShowAction keeps the Actions shown.
	ShowAction bool
	// Danger colors the item with the error color.
	Danger bool
	// Disabled dims the item and keeps it from being clicked.
	Disabled bool
	// Divider makes the item a line between groups; its other fields are ignored.
	Divider bool
}

type ListProps struct {
	// Items are the rows, in order.
	Items []ListItem
	// Active is the Key of the row shown as chosen.
	Active string
	// Variant picks the container: Filled tints it and Outlined borders it, both padded by 4.
	Variant ListVariant
	// Compact narrows the rows' side padding from 12 to 8.
	Compact bool
}

var listTransition = ui.ElementTransition{Duration: 150 * time.Millisecond, Ease: motion.EaseOut, Colors: true}

// List shows rows of navigation or settings, each with an icon, a label, an optional description
// and extra text; it returns the Key of the row clicked in the frame, or "".
func List(c *ui.Context, p ListProps) string {
	tk := theme.TokensOf(c)
	root := ui.Column(c).Gap(2).AlignSelf(ui.Stretch)
	switch p.Variant {
	case ListFilled:
		root.Padding(4).Radius(tk.BorderRadiusLG).Background(tk.FillQuaternary)
	case ListOutlined:
		root.Padding(4).Radius(tk.BorderRadiusLG).Border(1, tk.BorderSecondary)
	}
	clicked := ""
	root.Children(func() {
		for i, it := range p.Items {
			if it.Divider {
				ui.Box(c.Key(i)).Height(1).Margin(4, 0).AlignSelf(ui.Stretch).Background(tk.BorderSecondary)
				continue
			}
			if listRow(c.Key(it.Key), tk, it, it.Key == p.Active, p.Compact) {
				clicked = it.Key
			}
		}
	})
	return clicked
}

func listRow(c *ui.Context, tk *theme.Tokens, it ListItem, active, compact bool) bool {
	clicked := false
	cell := ui.Box(c).AlignSelf(ui.Stretch)
	cell.Children(func() {
		row := ui.ButtonBase(c).Label(it.Label).Disabled(it.Disabled).Row().Gap(12).AlignItems(ui.Center).
			Justify(ui.Start).MinHeight(32).PaddingY(6).PaddingX(12).Radius(tk.BorderRadius).
			FontSize(14).AlignSelf(ui.Stretch).Transition(listTransition)
		if compact {
			row.PaddingX(8)
		}
		fg := tk.TextSecondary
		switch {
		case it.Disabled:
			row.Opacity(0.45).Cursor(ui.CursorNotAllowed)
		case active:
			fg = tk.Text
			row.Background(tk.FillSecondary).Cursor(ui.CursorPointer)
		case row.Hovered():
			fg = tk.Text
			row.Background(tk.FillTertiary).Cursor(ui.CursorPointer)
		default:
			row.Cursor(ui.CursorPointer)
		}
		if it.Danger {
			fg = tk.Error
		}
		row.TextColor(fg)
		row.Children(func() {
			switch {
			case it.Avatar != nil:
				it.Avatar()
			case it.Icon != nil:
				ui.Icon(c, it.Icon).FontSize(16)
			}
			ui.Column(c).Gap(1).Grow(1).Basis(0).MinWidth(0).Children(func() {
				label := ui.Text(c, it.Label).SingleLine()
				if active {
					label.FontWeight(500)
				}
				if it.Description != "" {
					ui.Text(c, it.Description).SingleLine().FontSize(12).LineHeight(1.4).TextColor(tk.TextDescription)
				}
			})
			if it.Extra != "" {
				ui.Text(c, it.Extra).FontSize(12).TextColor(tk.TextDescription).Shrink(0)
			}
		})
		clicked = row.Clicked()
		if it.Actions != nil && (it.ShowAction || cell.Hovered() || cell.FocusWithin()) {
			ui.Row(c).Absolute().Right(8).Top(0).Bottom(0).Gap(4).AlignItems(ui.Center).Children(it.Actions)
		}
	})
	return clicked
}
