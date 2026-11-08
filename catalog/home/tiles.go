package home

import (
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/code"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

// Tile is a bento tile of the home page: a scene of real components, rendered to an image.
type Tile struct {
	Name          string
	Width, Height float32
	View          func(c *ui.Context)
}

var tileDay = time.Date(2026, time.November, 12, 0, 0, 0, 0, time.UTC)

const tileCode = `func App(c *ui.Context) {
	theme.Apply(c)
	ui.Column(c).Gap(12).Children(func() {
		display.Avatar(c, "🤖", display.AvatarProps{})
		action.Button(c, "Ship it", action.ButtonProps{
			Variant: action.Primary,
		})
	})
}`

var Tiles = []Tile{
	{"agents", 520, 260, agentsTile},
	{"code", 520, 260, func(c *ui.Context) {
		code.Highlighter(c, tileCode, code.HighlighterProps{Language: "go", FullFeatured: true, FileName: "main.go"}).Width(480)
	}},
	{"calendar", 260, 260, func(c *ui.Context) {
		day := tileDay
		field.Calendar(c, &day, field.CalendarProps{NoToday: true})
	}},
	{"qr", 260, 260, func(c *ui.Context) {
		code.QRCode(c, "https://mygo.lobehub.com", code.QRCodeProps{Size: 150})
	}},
	{"controls", 260, 260, controlsTile},
	{"buttons", 260, 120, func(c *ui.Context) {
		ui.Row(c).Gap(8).Children(func() {
			action.Button(c, "Primary", action.ButtonProps{Variant: action.Primary})
			action.Button(c, "Default", action.ButtonProps{})
		})
	}},
	{"group", 260, 120, func(c *ui.Context) {
		items := []display.GroupAvatarItem{}
		for _, e := range []string{"🦊", "🐳", "🦄", "🐙"} {
			items = append(items, display.GroupAvatarItem{Avatar: e})
		}
		display.GroupAvatar(c, display.GroupAvatarProps{Avatars: items, Size: 72})
	}},
	{"alerts", 520, 120, func(c *ui.Context) {
		ui.Column(c).Width(480).Gap(8).Children(func() {
			feedback.Alert(c, "Build finished: 3 platforms, 0 warnings", feedback.AlertProps{Type: feedback.Success})
			feedback.Alert(c, "macOS signing certificate expires in 12 days", feedback.AlertProps{Type: feedback.Warning})
		})
	}},
	{"swatches", 520, 120, swatchesTile},
}

func agentsTile(c *ui.Context) {
	shapes := []display.AgentShape{display.AgentFace, display.AgentCloud, display.AgentHex, display.AgentCat,
		display.AgentGhost, display.AgentStar, display.AgentPudding, display.AgentCapsule}
	states := []display.AgentState{display.AgentIdle, display.AgentThinking, display.AgentWorking, display.AgentDone}
	ui.Column(c).Gap(14).Children(func() {
		for r := range 2 {
			ui.Row(c.Key(r)).Gap(18).Children(func() {
				for i := range 4 {
					n := r*4 + i
					display.AgentAvatar(c.Key(n), display.AgentAvatarProps{Seed: string(rune('a' + n)), Shape: shapes[n],
						Color: display.AgentColor(n%7) + display.AgentPurple, Size: 88, State: states[(n+r)%4]})
				}
			})
		}
	})
}

func controlsTile(c *ui.Context) {
	on, checked := true, true
	level, stars := 64.0, 4.0
	tab := 1
	ui.Column(c).Width(220).Gap(16).Children(func() {
		action.Segmented(c, &tab, []string{"Chat", "Files", "Agents"}, action.SegmentedProps{Block: true})
		ui.Row(c).Gap(10).AlignItems(ui.Center).Children(func() {
			choice.Switch(c, &on, choice.SwitchProps{}).Label("Sync")
			ui.Text(c, "Sync across devices")
		})
		choice.Slider(c, &level, choice.SliderProps{}).AlignSelf(ui.Stretch)
		choice.Rate(c, &stars, choice.RateProps{})
		choice.Checkbox(c, &checked, "Remember me", choice.CheckboxProps{})
	})
}

func swatchesTile(c *ui.Context) {
	tk := theme.TokensOf(c)
	color := tk.Purple
	ui.Column(c).Gap(14).AlignItems(ui.Center).Children(func() {
		var swatches []choice.ColorSwatch
		for _, col := range []ui.Color{tk.Red, tk.Orange, tk.Gold, tk.Green, tk.Cyan, tk.Blue, tk.Geekblue, tk.Purple, tk.Magenta} {
			swatches = append(swatches, choice.ColorSwatch{Color: col})
		}
		choice.ColorSwatches(c, &color, choice.ColorSwatchesProps{Colors: swatches})
		ui.Row(c).Gap(6).Children(func() {
			for _, t := range []struct {
				label string
				color display.TagColor
			}{{"stable", display.TagGreen}, {"go 1.27", display.TagBlue}, {"no cgo", "purple"}, {"MIT", "orange"}} {
				display.Tag(c.Key(t.label), t.label, display.TagProps{Color: t.color})
			}
		})
	})
}
