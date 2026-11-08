package markdown

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/markdown"
)

const snippets = "With FullFeaturedCodeBlock every block gets the header: its language, Copy and a toggle that collapses it.\n\n" +
	"```ts\nexport const answer = (): number => 42;\n```\n\n" +
	"```py\ndef greet(name: str) -> str:\n    return f\"hello, {name}\"\n```\n"

func Code(c *ui.Context) {
	ui.Column(c).Width(560).Children(func() {
		markdown.Markdown(c, snippets, markdown.MarkdownProps{FullFeaturedCodeBlock: true})
	})
}
