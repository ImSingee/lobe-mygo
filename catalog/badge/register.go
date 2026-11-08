package badge

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("badge/basic", Basic)
	registry.Add("badge/status", Status)
}
