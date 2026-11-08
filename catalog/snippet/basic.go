package snippet

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/code"
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(12).AlignItems(ui.Start).Children(func() {
		code.Snippet(c.Key("filled"), "go get github.com/lobehub/lobe-mygo", code.SnippetProps{Language: "sh", Prefix: "$"})
		code.Snippet(c.Key("outlined"), "bun run dev", code.SnippetProps{Language: "sh", Prefix: "$", Variant: code.Outlined})
		code.Snippet(c.Key("borderless"), "npx wrangler deploy", code.SnippetProps{Language: "sh", Variant: code.Borderless})
	})
}
