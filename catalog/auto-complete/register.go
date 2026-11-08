package autocomplete

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("auto-complete/basic", Basic)
}
