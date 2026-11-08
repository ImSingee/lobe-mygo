package upload

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/field"
)

type picture struct {
	name string
	src  ui.ImageSource
}

func Dragger(c *ui.Context) {
	root := ui.Column(c).Gap(16).Width(420)
	pics := ui.Local(root, "pics", func() []picture { return nil })
	root.Children(func() {
		field.Upload(c, field.UploadProps{
			Dragger: true, Multiple: true, MaxCount: 6, Accept: "image/*",
			Description: "PNG, JPG or GIF, up to 6 at once",
			OnFiles: func(files []field.File) {
				for _, f := range files {
					data, err := f.ReadAll()
					if err != nil {
						continue
					}
					if img, err := ui.DecodeImage(data); err == nil {
						*pics = append(*pics, picture{f.Name, img})
					}
				}
			},
		}, nil)
		if len(*pics) > 0 {
			display.ImageGroup(c, func(g *display.ImageGallery) {
				for i, p := range *pics {
					display.Image(c.Key(i), p.src, display.ImageProps{Size: 64, Alt: p.name, Gallery: g})
				}
			}).Row().Wrap().Gap(8)
		}
	})
}
