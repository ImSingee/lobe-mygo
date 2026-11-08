package inputpassword

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

var key = "sk-lobe-1234567890"

func Basic(c *ui.Context) {
	ui.Column(c).Gap(12).Width(320).Children(func() {
		field.InputPassword(c, &key, field.InputProps{Prefix: lucide.KeyRound(), Placeholder: "API key"}).Label("API key")
	})
}
