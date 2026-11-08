package selectdemo

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var (
	lang  = []string{"en", "en", "en"}
	langs = []field.SelectOption{{Value: "en", Label: "English"}, {Value: "zh", Label: "简体中文"}, {Value: "ja", Label: "日本語"}}
)

func Variants(c *ui.Context) {
	ui.Column(c).Gap(12).Height(300).Children(func() {
		for i, v := range []field.InputVariant{field.InputOutlined, field.InputFilled, field.InputBorderless} {
			field.Select(c.Key(i), &lang[i], field.SelectProps{Options: langs, Variant: v}).Width(220)
		}
		field.Select(c.Key("small"), &lang[0], field.SelectProps{Options: langs, Size: theme.Small}).Width(220)
	})
}
