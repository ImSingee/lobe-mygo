package checkbox

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
)

var fruits = []struct {
	name    string
	checked bool
}{{"Apple", true}, {"Banana", false}, {"Cherry", false}}

func Group(c *ui.Context) {
	count := 0
	for _, f := range fruits {
		if f.checked {
			count++
		}
	}
	all := count == len(fruits)
	ui.Column(c).Gap(10).Children(func() {
		if choice.Checkbox(c, &all, "All fruits", choice.CheckboxProps{Indeterminate: count > 0 && count < len(fruits)}).Changed() {
			for i := range fruits {
				fruits[i].checked = all
			}
		}
		ui.Row(c).Gap(16).PaddingX(24).Children(func() {
			for i := range fruits {
				choice.Checkbox(c.Key(fruits[i].name), &fruits[i].checked, fruits[i].name, choice.CheckboxProps{})
			}
		})
	})
}
