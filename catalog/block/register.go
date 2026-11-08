package block

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("block/basic", Basic)
}
