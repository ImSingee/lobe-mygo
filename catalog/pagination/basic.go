package pagination

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/nav"
)

var (
	fewPage  = 1
	manyPage = 6
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).AlignItems(ui.Center).Children(func() {
		nav.Pagination(c.Key("few"), &fewPage, nav.PaginationProps{Total: 50})
		nav.Pagination(c.Key("many"), &manyPage, nav.PaginationProps{Total: 500})
	})
}
