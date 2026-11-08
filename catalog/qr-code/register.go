package qrcode

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("qr-code/basic", Basic)
	registry.Add("qr-code/status", Status)
}
