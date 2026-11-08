package progress

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("progress/basic", Basic)
	registry.Add("progress/variants", Variants)
}
