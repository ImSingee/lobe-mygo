package text

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("text/basic", Basic)
	registry.Add("text/headings", Headings)
	registry.Add("text/ellipsis", Ellipsis)
}
