package highlighter

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/code"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func FullFeatured(c *ui.Context) {
	ui.Column(c).Gap(16).Width(560).Children(func() {
		code.Highlighter(c.Key("go"), goSample, code.HighlighterProps{
			Language: "go", FullFeatured: true, FileName: "greet.go",
			Actions: func() {
				action.ActionIcon(c, lucide.Play(), action.ActionIconProps{Size: theme.Small, Title: "Run"})
			},
		})
		code.Highlighter(c.Key("ts"), tsSample, code.HighlighterProps{Language: "ts", FullFeatured: true, Wrap: true, DefaultCollapsed: true})
	})
}
