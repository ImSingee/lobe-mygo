package alert

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("alert/basic", Basic)
	registry.Add("alert/variants", Variants)
	registry.Add("alert/description", Description)
}
