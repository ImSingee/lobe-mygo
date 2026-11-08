package image

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Basic(c *ui.Context) {
	ui.Row(c).Gap(16).AlignItems(ui.Start).Children(func() {
		display.Image(c.Key("wave"), wave(), display.ImageProps{Width: 360, Height: 240, Alt: "The Great Wave off Kanagawa"})
		display.Image(c.Key("crest"), crest(), display.ImageProps{Size: 240, Alt: "The wave's crest"})
	})
}
