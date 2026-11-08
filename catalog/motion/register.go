package motion

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("motion/springs", Springs)
	registry.Add("motion/presets", Presets)
}
