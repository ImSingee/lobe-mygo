package empty

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("empty/basic", Basic)
	registry.Add("empty/page", Page)
}
