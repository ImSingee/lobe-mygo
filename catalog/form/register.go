package form

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("form/basic", Basic)
	registry.Add("form/validation", Validation)
	registry.Add("form/variants", Variants)
	registry.Add("form/autosave", Autosave)
}
