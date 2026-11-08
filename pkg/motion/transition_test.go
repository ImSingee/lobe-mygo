package motion

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func playTime(tt *ui.Tester) time.Duration {
	const step = 5 * time.Millisecond
	elapsed := time.Duration(0)
	for tt.FramePending() && elapsed < 3*time.Second {
		tt.Advance(step)
		elapsed += step
	}
	return elapsed
}

func enterExitTimes(t *testing.T, tr ui.ElementTransition) (enter, exit time.Duration) {
	t.Helper()
	shown := false
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Center().Children(func() {
			if shown {
				ui.Box(c).Key("panel").Size(120, 80).Background(ui.RGB(30, 30, 30)).Transition(tr)
			}
		})
	}, 300, 300)
	shown = true
	tt.Frame()
	enter = playTime(tt)
	shown = false
	tt.Frame()
	exit = playTime(tt)
	return enter, exit
}

// The tester's frames land on display refreshes, so the frame that ends an animation comes up to
// one refresh (about 17 ms) after its duration.
func near(got, want time.Duration) bool {
	d := got - want
	return d >= 0 && d <= 25*time.Millisecond
}

func TestPresetEnterAndExitDurations(t *testing.T) {
	cases := []struct {
		name        string
		tr          ui.ElementTransition
		enter, exit time.Duration
	}{
		{"popup", Popup(), 150 * time.Millisecond, 75 * time.Millisecond},
		{"modal", Modal(), 220 * time.Millisecond, 120 * time.Millisecond},
		{"backdrop", Backdrop(), 180 * time.Millisecond, 180 * time.Millisecond},
		{"drawer", Drawer(Right), 300 * time.Millisecond, 220 * time.Millisecond},
		{"collapse", Collapse(), 200 * time.Millisecond, 200 * time.Millisecond},
		{"toast", Toast(Bottom), 400 * time.Millisecond, 400 * time.Millisecond},
	}
	for _, c := range cases {
		enter, exit := enterExitTimes(t, c.tr)
		if !near(enter, c.enter) || !near(exit, c.exit) {
			t.Errorf("%s plays %v in and %v out, want %v and %v", c.name, enter, exit, c.enter, c.exit)
		}
	}
}

func TestTooltipExitsFast(t *testing.T) {
	_, exit := enterExitTimes(t, Tooltip())
	if !near(exit, 100*time.Millisecond) {
		t.Errorf("tooltip exits in %v, want 100ms", exit)
	}
}

func TestDrawerComesFromItsSide(t *testing.T) {
	for _, c := range []struct {
		side   Side
		xp, yp float32
	}{{Right, 100, 0}, {Left, -100, 0}, {Bottom, 0, 100}, {Top, 0, -100}} {
		m := Drawer(c.side).Enter
		if m.XPercent != c.xp || m.YPercent != c.yp || m.Opacity != 1 {
			t.Errorf("Drawer(%v).Enter = %+v", c.side, *m)
		}
	}
}
