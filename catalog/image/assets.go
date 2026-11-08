package image

import (
	_ "embed"
	"sync"

	"github.com/egoist/mygo/ui"
)

var (
	//go:embed wave.jpg
	waveJPG []byte
	//go:embed crest.jpg
	crestJPG []byte
	//go:embed fuji.jpg
	fujiJPG []byte
)

func decode(data []byte) func() ui.ImageSource {
	return sync.OnceValue(func() ui.ImageSource {
		img, err := ui.DecodeImage(data)
		if err != nil {
			return nil
		}
		return img
	})
}

var (
	wave  = decode(waveJPG)
	crest = decode(crestJPG)
	fuji  = decode(fujiJPG)
)
