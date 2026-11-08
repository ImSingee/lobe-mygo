package inputotp

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("input-otp/basic", Basic)
}
