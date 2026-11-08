package upload

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("upload/basic", Basic)
	registry.Add("upload/dragger", Dragger)
}
