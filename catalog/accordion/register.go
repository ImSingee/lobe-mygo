package accordion

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("accordion/basic", Basic)
	registry.Add("accordion/variants", Variants)
}
