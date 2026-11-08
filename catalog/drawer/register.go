package drawer

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("drawer/basic", Basic)
	registry.Add("drawer/placement", Placement)
}
