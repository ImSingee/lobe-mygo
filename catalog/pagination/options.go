package pagination

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/nav"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var optionPages = [4]int{3, 10, 2, 4}

func Options(c *ui.Context) {
	ui.Column(c).Gap(16).AlignItems(ui.Center).Children(func() {
		nav.Pagination(c.Key("total"), &optionPages[0], nav.PaginationProps{Total: 85, PageSize: 20,
			ShowTotal: func(total, from, to int) string { return fmt.Sprintf("%d-%d of %d items", from, to, total) }})
		nav.Pagination(c.Key("less"), &optionPages[1], nav.PaginationProps{Total: 500, ShowLessItems: true})
		nav.Pagination(c.Key("small"), &optionPages[2], nav.PaginationProps{Total: 120, Size: theme.Small})
		nav.Pagination(c.Key("disabled"), &optionPages[3], nav.PaginationProps{Total: 120, Disabled: true})
	})
}
