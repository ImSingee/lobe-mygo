package searchbar

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
)

var (
	query, filter string
	last          = "nothing yet"
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(12).Width(360).Children(func() {
		field.SearchBar(c, &query, field.SearchBarProps{
			ShortKey:    "mod+k",
			Placeholder: "Search agents...",
			OnSearch:    func(q string) { last = q },
		}).Label("Search agents")
		field.SearchBar(c, &filter, field.SearchBarProps{Variant: field.InputFilled, Loading: filter != ""}).Label("Filter")
		ui.Text(c, "Searched: "+last).FontSize(12)
	})
}
