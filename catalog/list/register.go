package list

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("list/basic", Basic)
	registry.Add("list/variants", Variants)
}
