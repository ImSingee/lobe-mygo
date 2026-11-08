package actionicon

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("action-icon/basic", Basic)
	registry.Add("action-icon/variants", Variants)
	registry.Add("action-icon/sizes", Sizes)
}
