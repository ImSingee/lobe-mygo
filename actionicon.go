package lobe

import "github.com/egoist/mygo/ui"

type ActionIconVariant int

const (
	ActionIconBorderless ActionIconVariant = iota
	ActionIconFilled
	ActionIconOutlined
)

type ActionIconProps struct {
	// Variant picks the surface: Borderless shows a fill only on hover, Filled keeps a neutral
	// fill, Outlined sits on the container with a border.
	Variant ActionIconVariant
	// Size sets the box (24, 36 or 44), corner radius and glyph size (14, 20 or 24).
	Size Size
	// Active keeps the hover fill, for a toggled or open state.
	Active bool
	// Danger colors the glyph with the error color.
	Danger bool
	// Title names the button for assistive technology and shows as its tooltip.
	Title string
}

func ActionIcon(c *ui.Context, svg *ui.SVG, p ActionIconProps) ui.Element {
	tk := TokensOf(c)
	box, radius, glyph := actionIconSize(p.Size)

	b := ui.ButtonBase(c).Size(box, box).Radius(radius).
		Cursor(ui.CursorPointer).Transition(buttonTransition)
	if p.Title != "" {
		b.Label(p.Title).Tooltip(p.Title)
	}

	state := stateOf(b)
	bg, border, _ := buttonColors(tk, actionIconSurface(p.Variant), state)
	if p.Active {
		bg = actionIconActiveFill(tk, p.Variant, state)
	}
	b.Background(bg).Border(1, border).TextColor(actionIconGlyphColor(tk, p.Danger, state))

	b.Children(func() { ui.Icon(c, svg).FontSize(glyph) })
	return b
}

func actionIconSize(s Size) (box, radius, glyph float32) {
	switch s {
	case Small:
		return 24, 4, 14
	case Large:
		return 44, 8, 24
	}
	return 36, 6, 20
}

func actionIconSurface(v ActionIconVariant) ButtonVariant {
	switch v {
	case ActionIconFilled:
		return Filled
	case ActionIconOutlined:
		return Default
	}
	return Text
}

func actionIconActiveFill(tk *Tokens, v ActionIconVariant, s buttonState) ui.Color {
	switch {
	case v == ActionIconFilled && s == buttonActive:
		return tk.Fill
	case v == ActionIconOutlined && s != buttonRest:
		return tk.Fill
	}
	return tk.FillSecondary
}

func actionIconGlyphColor(tk *Tokens, danger bool, s buttonState) ui.Color {
	if danger {
		switch s {
		case buttonActive:
			return tk.ErrorActive
		case buttonHover:
			return tk.Error
		}
		return tk.Error.Alpha(0.66)
	}
	switch s {
	case buttonActive:
		return tk.Text
	case buttonHover:
		return tk.TextSecondary
	}
	return tk.TextTertiary
}
