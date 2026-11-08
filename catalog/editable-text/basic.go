package editabletext

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
)

var (
	title   = "Weekly planning"
	editing bool
)

func Basic(c *ui.Context) {
	ui.Column(c).Width(360).Children(func() {
		field.EditableText(c, &title, &editing, field.EditableTextProps{})
	})
}
