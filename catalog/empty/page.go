package empty

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/display"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

func Page(c *ui.Context) {
	ui.Column(c).Width(560).Children(func() {
		display.Empty(c, display.EmptyProps{
			Page:        true,
			Icon:        lucide.FolderOpen(),
			Title:       "This folder is empty",
			Description: "Upload files or create a document to get started. Files you add are shared with everyone in the workspace.",
			Action: func() {
				action.Button(c, "Upload", action.ButtonProps{Variant: action.Primary})
				action.Button(c, "New document", action.ButtonProps{})
			},
		})
	})
}
