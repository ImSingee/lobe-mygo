package avatar

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("avatar/basic", Basic)
	registry.Add("avatar/fallback", Fallback)
	registry.Add("avatar/sizes", Sizes)
}
