package home

import (
	"strings"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/markdown"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

const streamStep = 12 * time.Millisecond

func chatPane(c *ui.Context, tk *theme.Tokens, st *showcase) {
	cv := st.convs[st.active]
	ui.Column(c).Grow(1).Basis(0).MinWidth(0).Children(func() {
		ui.Row(c).Height(56).Shrink(0).PaddingX(20).Gap(10).AlignItems(ui.Center).
			BorderWidth(0, 0, 1, 0).BorderColor(tk.BorderSecondary).Children(func() {
			agentAvatar(c.Key("top-"+cv.name), cv, 32, st.streaming >= 0)
			ui.Column(c).Grow(1).Basis(0).MinWidth(0).Children(func() {
				ui.Text(c, cv.name).FontWeight(600).SingleLine()
				ui.Text(c, cv.topic).FontSize(12).TextColor(tk.TextTertiary).SingleLine()
			})
			action.ActionIcon(c, lucide.Share2(), action.ActionIconProps{Title: "Share"})
			action.ActionIcon(c, lucide.Ellipsis(), action.ActionIconProps{Title: "More"})
		})
		ui.Scroll(c.Key("messages-" + cv.name)).Grow(1).Basis(0).Children(func() {
			ui.Column(c).AlignItems(ui.Center).Padding(24, 24, 8, 24).Children(func() {
				ui.Column(c).Width(720).MaxWidthPercent(100).Gap(24).Children(func() {
					for i, m := range cv.messages {
						messageRow(c.Key(i), tk, st, cv, i, m)
					}
				})
			})
		})
		composer(c, tk, st)
	})
}

func streamedText(c *ui.Context, st *showcase, text string) (string, bool) {
	n := int(c.Now().Sub(st.start) / streamStep)
	if n >= len(text) {
		return text, true
	}
	for n > 0 && n < len(text) && text[n]&0xC0 == 0x80 {
		n++
	}
	c.After(16 * time.Millisecond)
	return text[:n], false
}

func messageRow(c *ui.Context, tk *theme.Tokens, st *showcase, cv conversation, i int, m message) {
	if m.user {
		ui.Row(c).Justify(ui.End).Children(func() {
			ui.Box(c).Padding(10, 14).Radius(tk.BorderRadiusLG).Background(tk.FillTertiary).MaxWidthPercent(80).Children(func() {
				ui.Text(c, m.text).LineHeight(1.6)
			})
		})
		return
	}
	text, live := m.text, false
	if i == st.streaming {
		var done bool
		text, done = streamedText(c, st, m.text)
		live = !done
	}
	ui.Row(c).Gap(12).AlignItems(ui.Start).Children(func() {
		agentAvatar(c.Key("msg-"+cv.name), cv, 32, live)
		ui.Column(c).Grow(1).Basis(0).MinWidth(0).Padding(4, 0, 0, 0).Children(func() {
			md := markdown.Markdown(c, text, markdown.MarkdownProps{Variant: markdown.Chat, Animated: live})
			if live {
				md.ScrollIntoView()
			} else if i == st.streaming && markdown.Revealed(md) {
				st.streaming = -1
			}
		})
	})
}

func composer(c *ui.Context, tk *theme.Tokens, st *showcase) {
	ui.Column(c).Shrink(0).AlignItems(ui.Center).Padding(8, 24, 20, 24).Children(func() {
		box := ui.Column(c).Width(720).MaxWidthPercent(100).Padding(8).Gap(4).Radius(tk.BorderRadiusLG).
			Border(1, tk.BorderSecondary).Background(tk.BgContainer)
		tk.Elevate(box, theme.BoxShadowTertiary)
		box.Children(func() {
			field.TextArea(c, &st.draft, field.TextAreaProps{Variant: field.InputBorderless, MinRows: 2, MaxRows: 6,
				Placeholder: "Message " + st.convs[st.active].name + "…"}).Label("Message")
			ui.Row(c).Gap(4).AlignItems(ui.Center).Children(func() {
				action.ActionIcon(c, lucide.Paperclip(), action.ActionIconProps{Title: "Attach", Size: theme.Small})
				action.ActionIcon(c, lucide.Globe(), action.ActionIconProps{Title: "Search the web", Size: theme.Small})
				ui.Spacer(c)
				send := action.Button(c, "Send", action.ButtonProps{Variant: action.Primary, Size: theme.Small})
				if send.Clicked() && strings.TrimSpace(st.draft) != "" {
					sendDraft(c, st)
				}
			})
		})
	})
}

func sendDraft(c *ui.Context, st *showcase) {
	cv := &st.convs[st.active]
	cv.messages = append(cv.messages, message{user: true, text: strings.TrimSpace(st.draft)},
		message{text: replies[st.sent%len(replies)]})
	st.sent++
	st.draft = ""
	st.streaming = len(cv.messages) - 1
	st.start = c.Now()
}
