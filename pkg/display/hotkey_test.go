package display

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/internal/testkit"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func hotkeyTester(keys string, p HotkeyProps) *ui.Tester {
	return ui.NewTester(testkit.Centered(func(c *ui.Context) { Hotkey(c, keys, p) }), 300, 100)
}

func TestHotkeyNamesKeysForThePlatform(t *testing.T) {
	tt := hotkeyTester("mod+shift+k", HotkeyProps{Platform: HotkeyOther})
	for _, s := range []string{"Ctrl", "Shift", "K"} {
		if !tt.HasText(s) {
			t.Errorf("other platform: no %q in %q", s, tt.Texts())
		}
	}
	tt = hotkeyTester("mod+shift+k", HotkeyProps{Platform: HotkeyApple})
	if tt.HasText("Ctrl") || tt.HasText("Shift") || !tt.HasText("K") {
		t.Errorf("apple: texts %q, want glyphs for mod and shift and K", tt.Texts())
	}
	tt = hotkeyTester("enter", HotkeyProps{Platform: HotkeyOther})
	if !tt.HasText("Enter") {
		t.Errorf("enter on other platforms: %q", tt.Texts())
	}
}

func TestHotkeyChipsAndCompact(t *testing.T) {
	tk := theme.Light()
	tt := hotkeyTester("ctrl+k", HotkeyProps{Platform: HotkeyOther})
	ctrl, _ := tt.Find("Ctrl")
	k, _ := tt.Find("K")
	if gap := k.X - (ctrl.X + ctrl.W); gap < 2+16-0.5 {
		t.Errorf("separate chips: K %v DIPs after Ctrl, want 2 between chips plus 8 padding each side", gap)
	}
	if got := testkit.PixelAt(tt, 300, ctrl.X-4, ctrl.Y+ctrl.H/2); !testkit.Near(got, testkit.Over(tk.BgLayout, tk.FillTertiary)) {
		t.Errorf("filled chip %+v, want FillTertiary", got)
	}
	tt = hotkeyTester("ctrl+k", HotkeyProps{Platform: HotkeyOther, Compact: true})
	ctrl, _ = tt.Find("Ctrl")
	k, _ = tt.Find("K")
	if gap := k.X - (ctrl.X + ctrl.W); gap < 5.5 || gap > 6.5 {
		t.Errorf("compact: K %v DIPs after Ctrl, want 6 inside one chip", gap)
	}
	if got := testkit.PixelAt(tt, 300, ctrl.X+ctrl.W+3, ctrl.Y+ctrl.H/2); !testkit.Near(got, testkit.Over(tk.BgLayout, tk.FillTertiary)) {
		t.Errorf("compact: between keys %+v, want the one chip's fill", got)
	}
}
