package modal

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("modal/basic", Basic)
	registry.Add("modal/texts", Texts)
}
