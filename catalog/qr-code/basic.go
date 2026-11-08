package qrcode

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/code"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

func Basic(c *ui.Context) {
	ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
		code.QRCode(c.Key("plain"), "https://lobehub.com", code.QRCodeProps{})
		code.QRCode(c.Key("icon"), "https://lobehub.com", code.QRCodeProps{Icon: func(c *ui.Context) {
			display.Avatar(c, "🤯", display.AvatarProps{Size: 36, Background: ui.Hex("#ffffff")})
		}})
		code.QRCode(c.Key("color"), "https://mygo.lobehub.com", code.QRCodeProps{Size: 120,
			Color: ui.Hex("#5b3cc4"), Background: ui.Hex("#f4f0ff"), NoBorder: true})
	})
}
