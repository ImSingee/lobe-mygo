package skeleton

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("skeleton/basic", Basic)
	registry.Add("skeleton/animation", Animation)
}
