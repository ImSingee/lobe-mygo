package burger

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("burger/basic", Basic)
}
