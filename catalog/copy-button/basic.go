package copybutton

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func Basic(c *ui.Context) {
	ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
		ui.Text(c, "npm i @lobehub/ui").Font("monospace").FontSize(13)
		action.CopyButton(c, "npm i @lobehub/ui", action.CopyButtonProps{})
		action.CopyButton(c.Key("filled"), "npm i @lobehub/ui", action.CopyButtonProps{Variant: action.ActionIconFilled, Size: theme.Small})
	})
}
