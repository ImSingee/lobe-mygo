package home

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/catalog/registry"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func showcaseTester() *ui.Tester {
	return ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		Showcase(c)
	}, 1100, 640)
}

func TestShowcaseRegistered(t *testing.T) {
	if _, ok := registry.Get("home/showcase"); !ok {
		t.Error("home/showcase not registered")
	}
}

func TestShowcaseStreamsTheFirstReply(t *testing.T) {
	tt := showcaseTester()
	if !tt.HasText(conversations[0].messages[0].text) {
		t.Fatalf("texts %q, want the first question", tt.Texts())
	}
	tt.Advance(20 * time.Second)
	tt.Frame()
	if !tt.HasText("That is the whole app.") {
		t.Errorf("texts %q, want the reply streamed to its end", tt.Texts())
	}
}

func TestShowcaseSendsAMessage(t *testing.T) {
	tt := showcaseTester()
	tt.Advance(20 * time.Second)
	box, ok := tt.AccessByRole(ui.RoleTextField, "Message")
	if !ok {
		t.Fatal("no Message field")
	}
	tt.ClickAt(box.Bounds.X+20, box.Bounds.Y+box.Bounds.H/2)
	tt.Type("Ship it")
	send, _ := tt.AccessByRole(ui.RoleButton, "Send")
	tt.ClickAt(send.Bounds.X+send.Bounds.W/2, send.Bounds.Y+send.Bounds.H/2)
	tt.Advance(20 * time.Second)
	tt.Frame()
	if !tt.HasText("Ship it") {
		t.Errorf("texts %q, want the sent message", tt.Texts())
	}
	if !tt.HasText(replies[0][len(replies[0])-10:]) {
		t.Errorf("want the canned reply streamed in after the message")
	}
}

func TestShowcaseSwitchesConversations(t *testing.T) {
	tt := showcaseTester()
	second := conversations[1]
	row, ok := tt.AccessByRole(ui.RoleButton, second.name)
	if !ok {
		t.Fatalf("no %s row", second.name)
	}
	tt.ClickAt(row.Bounds.X+row.Bounds.W/2, row.Bounds.Y+row.Bounds.H/2)
	tt.Advance(20 * time.Second)
	tt.Frame()
	if !tt.HasText(second.messages[0].text) {
		t.Errorf("texts %q, want %s's conversation", tt.Texts(), second.name)
	}
}
