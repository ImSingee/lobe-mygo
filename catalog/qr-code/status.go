package qrcode

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/code"
)

func Status(c *ui.Context) {
	root := ui.Row(c).Gap(16)
	expired := ui.Local(root, "expired", func() bool { return true })
	root.Children(func() {
		code.QRCode(c.Key("loading"), "https://lobehub.com", code.QRCodeProps{Status: code.QRLoading})
		status := code.QRActive
		if *expired {
			status = code.QRExpired
		}
		code.QRCode(c.Key("expired"), "https://lobehub.com/login", code.QRCodeProps{Status: status,
			OnRefresh: func() { *expired = false }})
	})
}
