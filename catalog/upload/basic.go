package upload

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Column(c).Gap(12).Width(360)
	files := ui.Local(root, "files", func() []field.File { return nil })
	root.Children(func() {
		ui.Row(c).Children(func() {
			up := field.Upload(c, field.UploadProps{Multiple: true, OnFiles: func(f []field.File) { *files = append(*files, f...) }},
				func() {
					ui.Icon(c, lucide.Upload()).FontSize(14)
					ui.Text(c, "Choose files").SingleLine()
				})
			action.StyleButton(c, up, "", action.ButtonProps{})
		})
		for i, f := range *files {
			ui.Row(c.Key(i)).Gap(8).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, lucide.File()).FontSize(14).TextColor(tk.TextTertiary)
				ui.Text(c, f.Name).FontSize(13).SingleLine().Grow(1).MinWidth(0)
				ui.Text(c, sizeOf(f.Size)).FontSize(12).TextColor(tk.TextTertiary).FontFeatures("tnum")
				if action.ActionIcon(c, lucide.X(), action.ActionIconProps{Size: theme.Small}).Label("Remove " + f.Name).Clicked() {
					*files = append((*files)[:i:i], (*files)[i+1:]...)
				}
			})
		}
	})
}

func sizeOf(n int64) string {
	switch {
	case n >= 1<<20:
		return fmt.Sprintf("%.1f MB", float64(n)/(1<<20))
	case n >= 1<<10:
		return fmt.Sprintf("%.1f KB", float64(n)/(1<<10))
	}
	return fmt.Sprintf("%d B", n)
}
