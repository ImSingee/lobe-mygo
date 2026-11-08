package selectdemo

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

var model string

var models = []field.SelectOption{
	{Value: "opus", Label: "Claude Opus", Icon: lucide.Sparkles()},
	{Value: "sonnet", Label: "Claude Sonnet", Icon: lucide.Sparkles()},
	{Value: "haiku", Label: "Claude Haiku", Icon: lucide.Zap()},
	{Value: "legacy", Label: "Legacy model", Icon: lucide.Archive(), Disabled: true},
}

func Basic(c *ui.Context) {
	ui.Column(c).Height(260).Children(func() {
		field.Select(c, &model, field.SelectProps{Options: models, Placeholder: "Select a model"}).Width(260)
	})
}
