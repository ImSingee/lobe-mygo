package image

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("image/basic", Basic)
	registry.Add("image/states", States)
	registry.Add("image/gallery", Gallery)
}
