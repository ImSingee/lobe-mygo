package colorpicker

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("color-picker/basic", Basic)
	registry.Add("color-picker/triggers", Triggers)
}
