package switchdemo

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("switch/basic", Basic)
}
