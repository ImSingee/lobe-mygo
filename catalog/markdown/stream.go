package markdown

import (
	"strings"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/markdown"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var streamed = strings.Repeat(reply+"\n", 3) + article

type streamState struct {
	start     time.Time
	smoothing markdown.Smoothing
	byWord    bool
}

func Stream(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Column(c).Width(560).Gap(12)
	st := ui.Local(root, "stream", func() streamState { return streamState{} })
	if st.start.IsZero() {
		st.start = c.Now()
	}
	n := min(len(streamed), int(c.Now().Sub(st.start)/time.Millisecond)/4)
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
