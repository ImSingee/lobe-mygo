package floatingpanel

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var placements = []feedback.FloatingPanelPlacement{
	feedback.FloatingBottomRight, feedback.FloatingBottomLeft, feedback.FloatingTopRight, feedback.FloatingTopLeft,
}

var steps = []string{
	"Read the task and its linked issue",
	"Searched the codebase for the failing test",
	"Found the race in the upload queue",
	"Wrote a fix and a test that covers it",
	"Running the suite",
}

type state struct {
	open   bool
	corner int
	draft  string
}

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	root := ui.Column(c).Absolute().Left(0).Top(0).Right(0).Bottom(0).Padding(24).Gap(16)
	st := ui.Local(root, "panel", func() state { return state{open: true} })
	root.Children(func() {
		ui.Text(c, "Task detail workspace").FontSize(18).FontWeight(700)
		ui.Text(c, "Open a run panel beside the task; the page stays usable while it shows.").TextColor(tk.TextSecondary)
		ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
			if action.Button(c, "Open run panel", action.ButtonProps{Variant: action.Primary}).Clicked() {
				st.open = true
			}
			action.Segmented(c, &st.corner, []string{"Bottom right", "Bottom left", "Top right", "Top left"}, action.SegmentedProps{})
		})
	})
	feedback.FloatingPanel(c, &st.open, feedback.FloatingPanelProps{
		Title:     "Agent run",
		Placement: placements[st.corner],
		Width:     420,
		Actions: func() {
			action.ActionIcon(c, lucide.RotateCw(), action.ActionIconProps{Title: "Rerun", Size: theme.Small})
		},
		Footer: func() {
			ui.Box(c).Grow(1).Children(func() {
				field.Input(c, &st.draft, field.InputProps{Placeholder: "Ask a follow-up"})
			})
			action.Button(c, "Send", action.ButtonProps{Variant: action.Primary})
		},
	}, func() {
		ui.Column(c).Padding(16).Gap(12).Children(func() {
			for i, s := range steps {
				ui.Row(c.Key(i)).Gap(8).AlignItems(ui.Center).Children(func() {
					icon, color := lucide.CircleCheck(), tk.Success
					if i == len(steps)-1 {
						icon, color = lucide.LoaderCircle(), tk.TextTertiary
					}
					ui.Icon(c, icon).FontSize(16).TextColor(color)
					ui.Text(c, s)
				})
			}
		})
	})
}
