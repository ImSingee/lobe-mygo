package image

import (
	"strconv"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var downloads int

func Gallery(c *ui.Context) {
	tk := theme.TokensOf(c)
	pictures := []struct {
		src func() ui.ImageSource
		alt string
	}{{wave, "The Great Wave"}, {crest, "The crest"}, {fuji, "Mount Fuji"}}
	save := func(ui.ImageSource) { downloads++ }
	ui.Column(c).Gap(12).AlignItems(ui.Center).Children(func() {
		display.ImageGroup(c, func(g *display.ImageGallery) {
			for i, p := range pictures {
				display.Image(c.Key(i), p.src(), display.ImageProps{Width: 160, Height: 120, Alt: p.alt, Gallery: g, OnDownload: save})
			}
		}).Row().Gap(12)
		ui.Text(c, "Downloads asked for: "+strconv.Itoa(downloads)).FontSize(12).TextColor(tk.TextSecondary)
	})
}
