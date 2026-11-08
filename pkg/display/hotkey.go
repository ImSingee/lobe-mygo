package display

import (
	"strings"
	"sync"
	"unicode"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

type HotkeyVariant int

const (
	HotkeyFilled HotkeyVariant = iota
	HotkeyOutlined
	HotkeyBorderless
)

type HotkeyPlatform int

const (
	HotkeyAuto HotkeyPlatform = iota
	HotkeyApple
	HotkeyOther
)

type HotkeyProps struct {
	// Variant picks the key caps: Filled tints them, Outlined draws their border, Borderless shows
	// the keys as text in one run.
	Variant HotkeyVariant
	// Compact puts all the keys in one cap.
	Compact bool
	// Platform names the keys as Apple devices (⌘ ⌥ ⇧ glyphs) or as others do (Ctrl, Alt, Shift):
	// Auto asks the device the view runs on.
	Platform HotkeyPlatform
}

var isApple = sync.OnceValue(appleDevice)

// Hotkey shows a keyboard shortcut written as keys joined by "+", such as "mod+shift+k": mod is
// Command on Apple devices and Ctrl elsewhere.
func Hotkey(c *ui.Context, keys string, p HotkeyProps) ui.Element {
	tk := theme.TokensOf(c)
	apple := p.Platform == HotkeyApple || p.Platform == HotkeyAuto && isApple()
	parts := splitKeys(keys)
	joined := p.Compact || p.Variant == HotkeyBorderless
	gap := float32(2)
	if p.Variant == HotkeyBorderless {
		gap = 6
	}
	root := ui.Row(c).Gap(gap).AlignItems(ui.Center).Shrink(0)
	root.Children(func() {
		if joined {
			hotkeyCap(c, tk, p.Variant).Gap(6).Children(func() {
				for i, k := range parts {
					hotkeyKey(c.Key(i), k, apple)
				}
			})
			return
		}
		for i, k := range parts {
			hotkeyCap(c.Key(i), tk, p.Variant).Children(func() { hotkeyKey(c, k, apple) })
		}
	})
	return root
}

func hotkeyCap(c *ui.Context, tk *theme.Tokens, v HotkeyVariant) ui.Element {
	const font = 12
	e := ui.Row(c).AlignItems(ui.Center).Justify(ui.Center).Height(font * 1.8).MinWidth(font * 1.8).
		PaddingX(8).Radius(tk.BorderRadiusSM).FontSize(font).LineHeight(1.1).TextColor(tk.TextSecondary).Shrink(0)
	switch v {
	case HotkeyOutlined:
		e.Border(1, tk.BorderSecondary)
	case HotkeyBorderless:
		e.PaddingX(4)
	default:
		e.Background(tk.FillTertiary)
	}
	return e
}

type keyName struct {
	apple func() *ui.SVG
	scale float32
	text  string
	icon  func() *ui.SVG
}

var keyNames = map[string]keyName{
	"mod":       {apple: lucide.Command, scale: 0.95, text: "Ctrl"},
	"alt":       {apple: lucide.Option, scale: 0.95, text: "Alt"},
	"option":    {apple: lucide.Option, scale: 0.95, text: "Alt"},
	"ctrl":      {apple: lucide.ChevronUp, scale: 1, text: "Ctrl"},
	"control":   {apple: lucide.ChevronUp, scale: 1, text: "Ctrl"},
	"shift":     {apple: lucide.ArrowBigUp, scale: 1.15, text: "Shift"},
	"meta":      {apple: lucide.Command, scale: 0.95, icon: lucide.Grid2x2},
	"cmd":       {apple: lucide.Command, scale: 0.95, icon: lucide.Grid2x2},
	"command":   {apple: lucide.Command, scale: 0.95, icon: lucide.Grid2x2},
	"enter":     {apple: lucide.CornerDownLeft, scale: 1, text: "Enter"},
	"backspace": {apple: lucide.Delete, scale: 1, text: "Backspace"},
	"tab":       {apple: lucide.ArrowRightToLine, scale: 1, text: "Tab"},
	"up":        {icon: lucide.ArrowUp},
	"down":      {icon: lucide.ArrowDown},
	"left":      {icon: lucide.ArrowLeft},
	"right":     {icon: lucide.ArrowRight},
	"space":     {icon: lucide.Space},
	"comma":     {text: ","},
	"period":    {text: "."},
	"slash":     {text: "?"},
	"semicolon": {text: ";"},
	"minus":     {text: "-"},
	"equal":     {text: "+"},
}

func hotkeyKey(c *ui.Context, key string, apple bool) {
	n, ok := keyNames[strings.ToLower(key)]
	switch {
	case ok && apple && n.apple != nil:
		ui.Icon(c, n.apple()).FontSize(12 * n.scale)
	case ok && n.text != "":
		ui.Text(c, n.text).SingleLine()
	case ok && n.icon != nil:
		ui.Icon(c, n.icon()).FontSize(12)
	default:
		ui.Text(c, startCase(key)).SingleLine()
	}
}

func splitKeys(keys string) []string {
	var out []string
	for _, k := range strings.Split(strings.ReplaceAll(keys, "++", "+plus"), "+") {
		switch k = strings.TrimSpace(k); k {
		case "":
		case "plus":
			out = append(out, "+")
		default:
			out = append(out, k)
		}
	}
	return out
}

func startCase(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return s
	}
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
