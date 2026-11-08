package text

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

const paragraph = "Lobe UI is an open-source UI component library dedicated to building AIGC web applications. It inherits the design of LobeHub and keeps the components small, quiet and fast."

func Ellipsis(c *ui.Context) {
	ui.Column(c).Gap(16).Width(360).Children(func() {
		display.Text(c.Key(1), paragraph, display.TextProps{Rows: 1})
		display.Text(c.Key(2), paragraph, display.TextProps{Rows: 2, Type: display.TextSecondary})
	})
}
