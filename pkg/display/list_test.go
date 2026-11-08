package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/internal/testkit"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func listTester(active *string, clicked *[]string, actions *int) *ui.Tester {
	return ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		ui.Column(c).Padding(20).Width(320).Children(func() {
			k := List(c, ListProps{Active: *active, Items: []ListItem{
				{Key: "chat", Label: "Chat", Description: "Talk to agents", Icon: lucide.MessageSquare(), Extra: "12"},
				{Key: "files", Label: "Files", Actions: func() {
					if ui.ButtonBase(c).Label("Pin").Size(20, 20).Clicked() {
						*actions++
					}
				}},
				{Divider: true},
				{Key: "trash", Label: "Trash", Danger: true},
				{Key: "locked", Label: "Locked", Disabled: true},
			}})
			if k != "" {
				*clicked = append(*clicked, k)
				*active = k
			}
		})
	}, 360, 300)
}

func TestListClicksAndActive(t *testing.T) {
	tk := theme.Light()
	active, clicked, actions := "chat", []string{}, 0
	tt := listTester(&active, &clicked, &actions)
	chat, _ := tt.AccessByRole(ui.RoleButton, "Chat")
	if got := testkit.PixelAt(tt, 360, chat.Bounds.X+4, chat.Bounds.Y+chat.Bounds.H/2); !testkit.Near(got, testkit.Over(tk.BgLayout, tk.FillSecondary)) {
		t.Errorf("active row %+v, want FillSecondary", got)
	}
	if chat.Bounds.W < 280-0.5 {
		t.Errorf("row %v wide, want the list's width", chat.Bounds.W)
	}
	desc, ok := tt.Find("Talk to agents")
	if want := chat.Bounds.X + 12 + 16 + 12; !ok || desc.X != want || desc.Y < chat.Bounds.Y+chat.Bounds.H/2-1 {
		t.Errorf("description %v, want under the label at x %v past the 16 DIP icon", desc, want)
	}
	extra, _ := tt.Find("12")
	if extra.X+extra.W < chat.Bounds.X+chat.Bounds.W-12-0.5 {
		t.Errorf("extra %v not at the row end %v", extra, chat.Bounds)
	}
	tt.Click("Trash")
	tt.Frame()
	tt.Click("Locked")
	tt.Frame()
	if len(clicked) != 1 || clicked[0] != "trash" || active != "trash" {
		t.Errorf("clicked %q active %q", clicked, active)
	}
}

func TestListActionsShowOnHover(t *testing.T) {
	active, clicked, actions := "", []string{}, 0
	tt := listTester(&active, &clicked, &actions)
	if _, ok := tt.AccessByRole(ui.RoleButton, "Pin"); ok {
		t.Fatal("actions show while the row is not hovered")
	}
	files, _ := tt.AccessByRole(ui.RoleButton, "Files")
	tt.Move(files.Bounds.X+40, files.Bounds.Y+files.Bounds.H/2)
	tt.Advance(time.Second)
	tt.Frame()
	pin, ok := tt.AccessByRole(ui.RoleButton, "Pin")
	if !ok {
		t.Fatal("no actions on hover")
	}
	if end := files.Bounds.X + files.Bounds.W - 8; pin.Bounds.X+pin.Bounds.W < end-0.5 || pin.Bounds.X+pin.Bounds.W > end+0.5 {
		t.Errorf("actions end at %v, want 8 DIPs from the row end %v", pin.Bounds.X+pin.Bounds.W, end)
	}
	tt.ClickAt(pin.Bounds.X+10, pin.Bounds.Y+10)
	tt.Frame()
	if actions != 1 || len(clicked) != 0 {
		t.Errorf("action clicks %d, row clicks %q: want the action alone", actions, clicked)
	}
}
