package searchbar

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("search-bar/basic", Basic)
}
