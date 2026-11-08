package markdown

import (
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/markdown"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

const streamed = "Streaming fits a chat reply: text comes in at an **even pace**, each character fading in, and `half-written syntax` shows as it will most likely end.\n\n" +
	"- Lists grow item by item\n- Links like [LobeHub](https://lobehub.com) appear once complete\n- Code blocks grow line by line\n\n" +
	"```go\nfunc reply(w io.Writer, words []string) {\n\tfor _, w := range words {\n\t\tfmt.Fprint(w, w, \" \")\n\t}\n}\n```\n\n" +
	"| Preset | Catch-up |\n| --- | --- |\n| Realtime | 40 ms |\n| Balanced | 120 ms |\n| Silky | 170 ms |\n\n" +
	"That is the whole reply."

type streamState struct {
	start     time.Time
	smoothing markdown.Smoothing
	byWord    bool
}

func Stream(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Column(c).Width(560).Height(600).Gap(12)
	st := ui.Local(root, "stream", func() streamState { return streamState{} })
	if st.start.IsZero() {
		st.start = c.Now()
	}
	n := min(len(streamed), int(c.Now().Sub(st.start)/time.Millisecond)/12)
	for n < len(streamed) && n > 0 && streamed[n]&0xC0 == 0x80 {
		n++
	}
	done := n >= len(streamed)
	if !done {
		c.After(16 * time.Millisecond)
	}
	root.Children(func() {
		ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
			if action.Button(c, "Restart", action.ButtonProps{Size: theme.Small}).Clicked() {
				st.start = c.Now()
			}
			for i, name := range []string{"Balanced", "Realtime", "Silky"} {
				variant := action.Default
				if st.smoothing == markdown.Smoothing(i) {
					variant = action.Primary
				}
				if action.Button(c.Key(name), name, action.ButtonProps{Size: theme.Small, Variant: variant}).Clicked() {
					st.smoothing = markdown.Smoothing(i)
					st.start = c.Now()
				}
			}
		})
		bubble := ui.Column(c).Padding(12, 16).Radius(tk.BorderRadiusLG).Background(tk.FillQuaternary)
		bubble.Children(func() {
			md := markdown.Markdown(c, streamed[:n], markdown.MarkdownProps{Variant: markdown.Chat, Animated: !done, Smoothing: st.smoothing})
			if done && markdown.Revealed(md) {
				ui.Text(c, "Done").FontSize(12).TextColor(tk.TextTertiary)
			}
		})
	})
}
