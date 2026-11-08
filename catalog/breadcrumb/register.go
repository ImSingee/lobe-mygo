package breadcrumb

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("breadcrumb/basic", Basic)
}
