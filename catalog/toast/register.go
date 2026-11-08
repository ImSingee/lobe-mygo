package toast

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("toast/basic", Basic)
	registry.Add("toast/action", Action)
}
