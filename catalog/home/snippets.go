package home

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/markdown"
)

// Snippet is a code sample of the home page with the scene it renders, kept side by side so the
// code shown is the code that drew the preview.
type Snippet struct {
	Key, Label, Code string
	Tile             Tile
}

var helloCount = 3

const snippetReply = "Rendered by **Lobe MyGo**:\n\n- Lists and `inline code`\n- Tables and code blocks\n\n```go\nfmt.Println(\"hello\")\n```"

var Snippets = []Snippet{
	{Key: "button", Label: "Button", Code: `var (
	count   int
	primary = action.ButtonProps{Variant: action.Primary}
)

func Hello(c *ui.Context) {
	theme.Apply(c)
	ui.Column(c).Gap(12).Children(func() {
		if action.Button(c, "Hello", primary).Clicked() {
			count++
		}
		ui.Text(c, fmt.Sprintf("Clicked %d times", count))
	})
}`, Tile: Tile{"snippet-button", 480, 300, func(c *ui.Context) {
		ui.Column(c).Gap(12).AlignItems(ui.Center).Children(func() {
			action.Button(c, "Hello", action.ButtonProps{Variant: action.Primary})
			ui.Text(c, fmt.Sprintf("Clicked %d times", helloCount))
		})
	}}},
	{Key: "markdown", Label: "Markdown", Code: `func Reply(c *ui.Context) {
	theme.Apply(c)
	markdown.Markdown(c, reply, markdown.MarkdownProps{
		Variant:  markdown.Chat,
		Animated: streaming,
	})
}`, Tile: Tile{"snippet-markdown", 480, 300, func(c *ui.Context) {
		ui.Column(c).Width(400).Children(func() {
			markdown.Markdown(c, snippetReply, markdown.MarkdownProps{Variant: markdown.Chat})
		})
	}}},
	{Key: "agents", Label: "Agent avatar", Code: `var states = []display.AgentState{
	display.AgentIdle,
	display.AgentThinking,
	display.AgentDone,
}

func Agents(c *ui.Context) {
	theme.Apply(c)
	ui.Row(c).Gap(16).Children(func() {
		for _, s := range states {
			p := display.AgentAvatarProps{State: s}
			display.AgentAvatar(c.Key(s), p)
		}
	})
}`, Tile: Tile{"snippet-agents", 480, 300, func(c *ui.Context) {
		ui.Row(c).Gap(16).Children(func() {
			for _, s := range []display.AgentState{display.AgentIdle, display.AgentThinking, display.AgentDone} {
				display.AgentAvatar(c.Key(s), display.AgentAvatarProps{State: s, Size: 72})
			}
		})
	}}},
}
