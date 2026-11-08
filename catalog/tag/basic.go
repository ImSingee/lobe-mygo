package tag

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
	"github.com/lobehub/lobe-mygo/lucide"
)

var closed int

func Basic(c *ui.Context) {
	ui.Row(c).Gap(8).Children(func() {
		lobe.Tag(c, "Default", lobe.TagProps{})
		lobe.Tag(c, "With icon", lobe.TagProps{Icon: lucide.Sparkles(), Color: lobe.TagPurple})
		lobe.Tag(c, "Closable", lobe.TagProps{Closable: true, OnClose: func() { closed++ }})
	})
	ui.Text(c, fmt.Sprintf("Closed %d tags", closed)).TextColor(lobe.TokensOf(c).TextSecondary)
}
