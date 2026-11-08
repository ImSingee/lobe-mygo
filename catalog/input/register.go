package input

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("input/basic", Basic)
	registry.Add("input/variants", Variants)
	registry.Add("input/sizes", Sizes)
}
