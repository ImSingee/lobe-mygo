package divider

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("divider/basic", Basic)
	registry.Add("divider/vertical", Vertical)
}
