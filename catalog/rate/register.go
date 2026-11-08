package rate

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("rate/basic", Basic)
	registry.Add("rate/half", Half)
}
