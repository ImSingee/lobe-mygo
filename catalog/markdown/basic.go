package markdown

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/markdown"
)

const article = "# Lobe MyGo\n\n" +
	"Lobe UI for [mygo](https://github.com/egoist/mygo): **bold**, *italic*, ~~struck~~, `inline code` and a footnote[^1].\n\n" +
	"## Lists\n\n- Immediate-mode views\n- Native and WebAssembly\n  - macOS, Windows, Linux\n  - Browsers\n\n" +
	"1. Write the view\n2. Run it\n\n- [x] Highlighter\n- [ ] Mermaid\n\n" +
	"> Quotes keep a bar along their side and the secondary text color.\n\n" +
	"> [!TIP]\n> GitHub's alerts take Lobe's status colors.\n\n" +
	"> [!WARNING]\n> Pickers are niche; common components come first.\n\n" +
	"### Table\n\n| Prop | Type | Default |\n| --- | --- | --- |\n| `Variant` | `Variant` | `Default` |\n| `FontSize` | `float32` | 16 |\n\n" +
	"---\n\n[^1]: Footnotes gather at the end.\n"

func Basic(c *ui.Context) {
	ui.Column(c).Width(640).Children(func() {
		markdown.Markdown(c, article, markdown.MarkdownProps{})
	})
}
