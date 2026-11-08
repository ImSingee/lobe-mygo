package overlay

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/internal/testkit"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func modalTester(open *bool, p ModalProps, oks *int) *ui.Tester {
	return ui.NewTester(testkit.Centered(func(c *ui.Context) {
		if Modal(c, open, p, func() { ui.Text(c, "Delete 3 files?") }) {
			*oks++
		}
	}), 800, 600)
}

// The panel's pose runs in frames that only paint, and Find reads the boxes of the last build, so a
// frame is built once the pose has settled.
func settle(tt *ui.Tester) {
	tt.Advance(time.Second)
	tt.Frame()
}

func TestModalCentersItsPanelOverTheBackdrop(t *testing.T) {
	tk := theme.Light()
	open, oks := true, 0
	tt := modalTester(&open, ModalProps{Title: "Delete files", Width: 400}, &oks)
	settle(tt)
	for _, s := range []string{"Delete files", "Delete 3 files?", "OK", "Cancel"} {
		if !tt.HasText(s) {
			t.Errorf("no %q in %v", s, tt.Texts())
		}
	}
	ok, _ := tt.Find("OK")
	if right := ok.X + ok.W; right < 600-16-15 || right > 600-16 {
		t.Errorf("OK ends at %v, want at the panel's end 584 less the button's padding", right)
	}
	title, _ := tt.Find("Delete files")
	if title.X < 200+16-1 || title.X > 200+16+1 {
		t.Errorf("title at x %v, want 216", title.X)
	}
	want := testkit.Over(tk.BgLayout, tk.BgContainer.Alpha(0.6))
	if got := testkit.PixelAt(tt, 800, 5, 5); !testkit.Near(got, want) {
		t.Errorf("backdrop = %+v, want BgContainer at 60%% %+v", got, want)
	}
	if got := testkit.PixelAt(tt, 800, 210, title.Y+title.H/2); !testkit.Near(got, tk.BgElevated) {
		t.Errorf("panel = %+v, want BgElevated", got)
	}
}

func TestModalOkReportsWithoutClosingAndCancelCloses(t *testing.T) {
	open, oks := true, 0
	tt := modalTester(&open, ModalProps{Title: "Delete files"}, &oks)
	tt.Advance(time.Second)
	clickText(t, tt, "OK")
	if oks != 1 || !open {
		t.Fatalf("after OK: oks %d open %v, want 1 and still open", oks, open)
	}
	clickText(t, tt, "Cancel")
	if open || oks != 1 {
		t.Errorf("after Cancel: open %v oks %d, want closed and no OK", open, oks)
	}
}

func TestModalDismissesWithEscapeBackdropAndClose(t *testing.T) {
	for name, dismiss := range map[string]func(tt *ui.Tester){
		"escape":   func(tt *ui.Tester) { tt.Key(0, ui.KeyEscape) },
		"backdrop": func(tt *ui.Tester) { tt.ClickAt(10, 10) },
		"close": func(tt *ui.Tester) {
			info, ok := tt.AccessByRole(ui.RoleButton, "Close")
			if !ok {
				t.Fatal("no Close button")
			}
			tt.ClickAt(info.Bounds.X+info.Bounds.W/2, info.Bounds.Y+info.Bounds.H/2)
		},
	} {
		open, oks := true, 0
		tt := modalTester(&open, ModalProps{Title: "Delete files"}, &oks)
		tt.Advance(time.Second)
		dismiss(tt)
		tt.Advance(time.Second)
		if open {
			t.Errorf("%s did not close the modal", name)
		}
	}
}

func TestModalTextsAndNoFooter(t *testing.T) {
	open, oks := true, 0
	tt := modalTester(&open, ModalProps{OkText: "Delete", CancelText: "Keep"}, &oks)
	tt.Advance(time.Second)
	if !tt.HasText("Delete") || !tt.HasText("Keep") || tt.HasText("OK") {
		t.Errorf("texts = %v, want Delete and Keep in place of OK and Cancel", tt.Texts())
	}
	tt = modalTester(&open, ModalProps{NoFooter: true}, &oks)
	tt.Advance(time.Second)
	if tt.HasText("OK") || tt.HasText("Cancel") {
		t.Errorf("NoFooter still shows buttons: %v", tt.Texts())
	}
}

func darkest(tt *ui.Tester, width int, r ui.Rect) int {
	low := 255 * 3
	for x := r.X; x < r.X+r.W; x++ {
		for y := r.Y; y < r.Y+r.H; y++ {
			p := testkit.PixelAt(tt, width, x, y)
			low = min(low, int(p.R)+int(p.G)+int(p.B))
		}
	}
	return low
}

func TestModalAnimatesOut(t *testing.T) {
	open, oks := true, 0
	tt := modalTester(&open, ModalProps{Title: "Delete files"}, &oks)
	tt.Advance(time.Second)
	title, _ := tt.Find("Delete files")
	tt.Key(0, ui.KeyEscape)
	tt.Advance(30 * time.Millisecond)
	if low := darkest(tt, 800, title); low > 3*160 {
		t.Errorf("title gone at once (darkest %d), want the dialog animating out", low)
	}
	tt.Advance(time.Second)
	if low := darkest(tt, 800, title); low < 3*230 {
		t.Errorf("title still painted (darkest %d) after the exit", low)
	}
}

func leftInk(tt *ui.Tester, width int, r ui.Rect) float32 {
	for x := r.X - 6; x < r.X+r.W; x += 0.5 {
		for y := r.Y; y < r.Y+r.H; y++ {
			p := testkit.PixelAt(tt, width, x, y)
			if int(p.R)+int(p.G)+int(p.B) < 3*235 {
				return x
			}
		}
	}
	return -1
}

func TestModalPanelScalesInFromTheCenter(t *testing.T) {
	open, oks := true, 0
	tt := modalTester(&open, ModalProps{Title: "Delete files"}, &oks)
	settle(tt)
	title, _ := tt.Find("Delete files")
	rest := leftInk(tt, 800, title)
	for _, pass := range []string{"first", "again"} {
		open = false
		settle(tt)
		open = true
		tt.Frame()
		tt.Advance(30 * time.Millisecond)
		if start := leftInk(tt, 800, title); start-rest < 3 {
			t.Errorf("%s open: title ink at x %v, rest %v: want the panel scaled toward its center", pass, start, rest)
		}
		settle(tt)
		if end := leftInk(tt, 800, title); end != rest {
			t.Errorf("%s open: title ink settles at %v, want %v", pass, end, rest)
		}
	}
}
