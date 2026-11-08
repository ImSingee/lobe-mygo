package home

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func ogTester() *ui.Tester {
	tt := ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		OG(c)
	}, 1280, 640)
	tt.SetDark(true)
	tt.Frame()
	return tt
}

func TestOGLeftColumn(t *testing.T) {
	tt := ogTester()
	name, ok := tt.Find("lobe-mygo")
	if !ok || name.X > 260 || name.Y > 160 {
		t.Errorf("name %+v, want at the top left after the marks", name)
	}
	first, ok1 := tt.Find(ogHeadline[0])
	second, ok2 := tt.Find(ogHeadline[1])
	if !ok1 || !ok2 || first.X != 64 || second.Y <= first.Y || first.H < 60 {
		t.Errorf("headline %+v / %+v, want two large lines at x 64", first, second)
	}
	desc, ok := tt.Find(ogDescription[0])
	if !ok || desc.Y < second.Y+second.H+40 || desc.X != 64 {
		t.Errorf("description %+v, want under the headline at x 64", desc)
	}
	if right := first.X + first.W; right > 680 {
		t.Errorf("headline reaches x %v, want it clear of the mascots", right)
	}
}

func TestOGArtFillsTheCard(t *testing.T) {
	tt := ogTester()
	art, ok := tt.Find("art")
	if !ok || art.X != 0 || art.Y != 0 || art.W != 1280 || art.H != 640 {
		t.Errorf("art %+v, want the illustration filling the card", art)
	}
}
