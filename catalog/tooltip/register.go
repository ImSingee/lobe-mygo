package tooltip

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("tooltip/basic", Basic)
	registry.Add("tooltip/placement", Placement)
}
