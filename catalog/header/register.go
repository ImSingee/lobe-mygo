package header

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("header/basic", Basic)
}
