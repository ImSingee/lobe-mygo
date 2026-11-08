package scrollarea

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("scroll-area/basic", Basic)
	registry.Add("scroll-area/horizontal", Horizontal)
}
