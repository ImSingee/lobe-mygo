package imageselect

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("image-select/basic", Basic)
}
