package tag

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("tag/basic", Basic)
	registry.Add("tag/colors", Colors)
	registry.Add("tag/variants", Variants)
	registry.Add("tag/sizes", Sizes)
}
