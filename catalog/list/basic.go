package list

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var active = "translator"

func Basic(c *ui.Context) {
	actions := func() {
		action.ActionIcon(c, lucide.Pin(), action.ActionIconProps{Size: theme.Small, Title: "Pin"})
		action.ActionIcon(c, lucide.Ellipsis(), action.ActionIconProps{Size: theme.Small, Title: "More"})
	}
	avatar := func(text string) func() {
		return func() { display.Avatar(c, text, display.AvatarProps{Size: 32, Shape: display.AvatarCircle}) }
	}
	ui.Column(c).Width(320).Children(func() {
		k := display.List(c, display.ListProps{Active: active, Items: []display.ListItem{
			{Key: "translator", Label: "Translator", Description: "Translate anything into 40 languages", Avatar: avatar("Translator"), Actions: actions},
			{Key: "writer", Label: "Writer", Description: "Drafts, edits and polishes", Avatar: avatar("Writer"), Actions: actions},
			{Key: "coder", Label: "Coder", Description: "Reads and writes code with you", Avatar: avatar("Coder"), Actions: actions},
		}})
		if k != "" {
			active = k
		}
	})
}
