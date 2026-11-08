package registry

import (
	"slices"

	"github.com/egoist/mygo/ui"
)

var demos = map[string]func(c *ui.Context){}

func Add(name string, view func(c *ui.Context)) {
	if _, ok := demos[name]; ok {
		panic("catalog: two demos named " + name)
	}
	demos[name] = view
}

func Get(name string) (func(c *ui.Context), bool) {
	view, ok := demos[name]
	return view, ok
}

func Names() []string {
	names := make([]string, 0, len(demos))
	for name := range demos {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}
