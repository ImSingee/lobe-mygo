package steps

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("steps/basic", Basic)
	registry.Add("steps/vertical", Vertical)
}
