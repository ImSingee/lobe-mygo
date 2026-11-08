package segmented

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("segmented/basic", Basic)
	registry.Add("segmented/variants", Variants)
}
