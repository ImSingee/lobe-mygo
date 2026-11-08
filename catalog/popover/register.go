package popover

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("popover/basic", Basic)
	registry.Add("popover/placement", Placement)
}
