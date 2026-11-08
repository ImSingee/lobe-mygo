package spin

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("spin/basic", Basic)
	registry.Add("spin/content", Content)
}
