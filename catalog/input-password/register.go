package inputpassword

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("input-password/basic", Basic)
}
