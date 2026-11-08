package pagination

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("pagination/basic", Basic)
	registry.Add("pagination/options", Options)
}
