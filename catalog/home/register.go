package home

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("home/showcase", Showcase)
}
