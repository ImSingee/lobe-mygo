package draggablepanel

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("draggable-panel/basic", Basic)
	registry.Add("draggable-panel/placement", Placement)
}
