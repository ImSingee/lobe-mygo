package autocomplete

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

var (
	model   string
	options = []field.AutoCompleteOption{
		{Value: "claude-opus-4", Label: "Claude Opus 4"},
		{Value: "claude-sonnet-4", Label: "Claude Sonnet 4"},
		{Value: "gpt-5", Label: "GPT-5"},
		{Value: "gemini-2.5-pro", Label: "Gemini 2.5 Pro"},
		{Value: "deepseek-r1", Label: "DeepSeek R1", Disabled: true},
	}
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(8).Width(320).Height(280).Children(func() {
		field.AutoComplete(c, &model, field.AutoCompleteProps{
			Options: options, EmptyText: "No models found", Placeholder: "Type a model name", Prefix: lucide.Bot(),
		}).Label("Model")
		ui.Text(c, "Model: "+model).FontSize(12)
	})
}
