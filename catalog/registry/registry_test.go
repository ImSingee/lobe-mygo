package registry

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"
)

func TestAddGetNames(t *testing.T) {
	defer reset()
	view := func(c *ui.Context) {}
	Add("b/two", view)
	Add("a/one", view)
	if _, ok := Get("a/one"); !ok {
		t.Error("Get(a/one) not found")
	}
	if _, ok := Get("zzz"); ok {
		t.Error("Get(zzz) found")
	}
	if got := Names(); !slices.Equal(got, []string{"a/one", "b/two"}) {
		t.Errorf("Names = %v", got)
	}
}

func TestAddDuplicatePanics(t *testing.T) {
	defer reset()
	Add("x", func(c *ui.Context) {})
	defer func() {
		if recover() == nil {
			t.Error("second Add of x did not panic")
		}
	}()
	Add("x", func(c *ui.Context) {})
}

func reset() { demos = map[string]func(c *ui.Context){} }
