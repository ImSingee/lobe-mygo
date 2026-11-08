package highlighter

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/code"
)

func Variants(c *ui.Context) {
	ui.Column(c).Gap(16).Width(560).Children(func() {
		for i, v := range []code.Variant{code.Filled, code.Outlined, code.Borderless} {
			code.Highlighter(c.Key(i), `const answer = 42; // the answer`, code.HighlighterProps{Language: "js", Variant: v})
		}
		code.Highlighter(c.Key("shadow"), `const answer = 42; // with a shadow`, code.HighlighterProps{Language: "js", Variant: code.Outlined, Shadow: true})
	})
}
