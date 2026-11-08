package togglegroup

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("toggle-group/basic", Basic)
	registry.Add("toggle-group/variants", Variants)
}
