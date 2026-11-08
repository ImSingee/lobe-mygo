package text

import (
	"strconv"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Headings(c *ui.Context) {
	ui.Column(c).Gap(4).Children(func() {
		for h := 1; h <= 5; h++ {
			display.Text(c.Key(h), "h"+strconv.Itoa(h)+". Lobe UI", display.TextProps{Heading: h})
		}
	})
}
