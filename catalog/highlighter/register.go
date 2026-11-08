package highlighter

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("highlighter/basic", Basic)
	registry.Add("highlighter/variants", Variants)
	registry.Add("highlighter/full-featured", FullFeatured)
}
