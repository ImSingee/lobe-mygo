package tokens

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("tokens/swatches", Swatches)
}
