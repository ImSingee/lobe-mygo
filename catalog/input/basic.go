package input

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

var query, email string

func Basic(c *ui.Context) {
	ui.Column(c).Gap(12).Width(320).Children(func() {
		field.Input(c, &query, field.InputProps{Placeholder: "Search", Prefix: lucide.Search()})
		field.Input(c, &email, field.InputProps{Placeholder: "Email", Invalid: email != "" && !validEmail(email)})
	})
}

func validEmail(s string) bool {
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '@' {
			return true
		}
	}
	return false
}
