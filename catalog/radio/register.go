package radio

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("radio/basic", Basic)
	registry.Add("radio/vertical", Vertical)
}
