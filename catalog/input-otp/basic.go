package inputotp

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
)

var (
	code     string
	invite   string
	verified = "Enter the 6-digit code"
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Children(func() {
		field.InputOTP(c, &code, field.InputOTPProps{OnComplete: func(s string) { verified = "Verified " + s }}).Label("Code")
		ui.Text(c, verified).FontSize(12)
		field.InputOTP(c, &invite, field.InputOTPProps{Length: 4, Alphanumeric: true, Variant: field.InputFilled}).Label("Invite")
	})
}
