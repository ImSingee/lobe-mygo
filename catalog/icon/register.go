package icon

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("icon/sizes", Sizes)
	registry.Add("icon/spin", Spin)
}
