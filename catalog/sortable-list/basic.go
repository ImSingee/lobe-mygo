package sortablelist

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/layout"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

type task struct {
	id, name string
}

type state struct {
	tasks   []task
	variant int
}

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Column(c).Width(360).Gap(16)
	st := ui.Local(root, "sortable", func() state {
		return state{tasks: []task{
			{"1", "Write the release notes"},
			{"2", "Review the open pull requests"},
			{"3", "Update the design tokens"},
			{"4", "Record the demo video"},
		}}
	})
	root.Children(func() {
		ui.Row(c).Children(func() {
			action.Segmented(c, &st.variant, []string{"Borderless", "Filled", "Outlined"}, action.SegmentedProps{})
		})
		layout.SortableList(c, &st.tasks, func(t task) string { return t.id },
			layout.SortableListProps{Gap: 4, Variant: layout.SortableVariant(st.variant)},
			func(t task, handle func()) {
				handle()
				ui.Text(c, t.name).Grow(1).SingleLine()
			})
		order := ""
		for i, t := range st.tasks {
			if i > 0 {
				order += " · "
			}
			order += t.id
		}
		ui.Text(c, "Order: "+order).FontSize(12).TextColor(tk.TextTertiary)
	})
}
