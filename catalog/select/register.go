package selectdemo

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("select/basic", Basic)
	registry.Add("select/variants", Variants)
}
