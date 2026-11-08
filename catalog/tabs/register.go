package tabs

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("tabs/basic", Basic)
	registry.Add("tabs/variants", Variants)
	registry.Add("tabs/vertical", Vertical)
}
