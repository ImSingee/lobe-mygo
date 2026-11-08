package snippet

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("snippet/basic", Basic)
}
