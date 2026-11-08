package floatingsheet

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("floating-sheet/basic", Basic)
	registry.Add("floating-sheet/inline", Inline)
}
