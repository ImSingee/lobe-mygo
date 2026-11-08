package floatingsheet

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func notes(c *ui.Context, tk *theme.Tokens) {
	ui.Column(c).Padding(4, 16, 16, 16).Gap(10).Children(func() {
		for i := range 12 {
			ui.Text(c.Key(i), fmt.Sprintf("Note %d: drag the header to resize, or down to close.", i+1)).TextColor(tk.TextSecondary)
		}
	})
}

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Column(c).Absolute().Left(0).Top(0).Right(0).Bottom(0).Background(tk.BgLayout)
	open := ui.Local(root, "open", func() bool { return true })
	root.Children(func() {
		ui.Column(c).Padding(24).Gap(12).AlignItems(ui.Start).Children(func() {
			ui.Text(c, "Conversation").FontSize(18).FontWeight(700)
			ui.Text(c, "The sheet floats over this content and snaps to 30%, 60% and 90% of the container.").TextColor(tk.TextSecondary)
			if action.Button(c, "Open sheet", action.ButtonProps{Variant: action.Primary}).Clicked() {
				*open = true
			}
		})
		feedback.FloatingSheet(c, open, feedback.FloatingSheetProps{
			Title:      "Notes",
			SnapPoints: []float32{0.3, 0.6, 0.9},
			MinHeight:  120,
			HeaderActions: func() {
				if action.ActionIcon(c, lucide.X(), action.ActionIconProps{Title: "Close", Size: theme.Small}).Clicked() {
					*open = false
				}
			},
		}, func() { notes(c, tk) })
	})
}

func Inline(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Column(c).Absolute().Left(0).Top(0).Right(0).Bottom(0).Padding(16).Gap(8).Background(tk.BgLayout)
	open := ui.Local(root, "open", func() bool { return true })
	root.Children(func() {
		ui.Column(c).Grow(1).Basis(0).Padding(8).Gap(8).AlignItems(ui.Start).Children(func() {
			ui.Text(c, "Inline mode keeps room for the sheet; dragged taller, it covers the content above.").TextColor(tk.TextSecondary)
			if action.Button(c, "Open sheet", action.ButtonProps{}).Clicked() {
				*open = true
			}
		})
		feedback.FloatingSheet(c, open, feedback.FloatingSheetProps{
			Title:         "Details",
			Inline:        true,
			Embedded:      true,
			RestingHeight: 160,
			MinHeight:     120,
		}, func() { notes(c, tk) })
	})
}
