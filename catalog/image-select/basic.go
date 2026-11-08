package imageselect

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

func Basic(c *ui.Context) {
	root := ui.Column(c)
	mode := ui.Local(root, "mode", func() string { return "auto" })
	root.Children(func() {
		choice.ImageSelect(c, mode, choice.ImageSelectProps[string]{Options: []choice.ImageSelectItem[string]{
			{Value: "light", Label: "Light", Icon: lucide.Sun, Image: lightPreview()},
			{Value: "dark", Label: "Dark", Icon: lucide.Moon, Image: darkPreview()},
			{Value: "auto", Label: "System", Icon: lucide.Monitor, Image: autoPreview()},
		}})
	})
}
