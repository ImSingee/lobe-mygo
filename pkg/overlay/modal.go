package overlay

import (
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
	"github.com/lobehub/lobe-mygo/pkg/motion"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

type ModalProps struct {
	// Title shows in the header; without one, the content starts at the top.
	Title string
	// Width is the panel's width, 520 DIPs when zero, and at most the window's less 32.
	Width float32
	// OkText labels the primary button, "OK" when empty.
	OkText string
	// CancelText labels the other button, "Cancel" when empty.
	CancelText string
	// NoFooter hides both buttons.
	NoFooter bool
}

var closeTransition = ui.ElementTransition{Duration: 160 * time.Millisecond, Ease: motion.SoftEase, Colors: true}

// Modal shows content in a dialog over the window while *open is true. Cancel, the close button,
// Escape and a click on the backdrop close it. It returns true in the frame OK is clicked and
// leaves *open to the caller, which closes it once the action succeeds.
func Modal(c *ui.Context, open *bool, p ModalProps, content func()) bool {
	tk := theme.TokensOf(c)
	ok := false
	ui.DialogBase(c, open, func(backdrop, panel ui.Element) {
		backdrop.Padding(32, 16).Background(tk.BgContainer.Alpha(0.6)).Transition(motion.Backdrop())
		width := p.Width
		if width == 0 {
			width = 520
		}
		panel.Width(width).MaxWidthPercent(100).MaxHeightPercent(100).Radius(12).Clip().
			Background(tk.BgElevated).TextColor(tk.Text).FontSize(tk.FontSize)
		scaleIn(c, panel)
		tk.Elevate(panel, theme.BoxShadow)
		panel.Children(func() {
			if p.Title != "" {
				ui.Row(c).MinHeight(48).Padding(8, 56, 8, 16).Children(func() {
					ui.Text(c, p.Title).FontSize(14).FontWeight(600)
				})
			}
			body := ui.Scroll(c).Shrink(1)
			body.Children(func() {
				col := ui.Column(c).Padding(0, 16, 16, 16)
				if p.Title == "" {
					col.Padding(16)
				}
				col.Children(content)
			})
			if !p.NoFooter {
				ui.Row(c).Gap(8).Padding(12, 16).Justify(ui.End).Children(func() {
					if action.Button(c, textOr(p.CancelText, "Cancel"), action.ButtonProps{}).Clicked() {
						*open = false
					}
					if action.Button(c, textOr(p.OkText, "OK"), action.ButtonProps{Variant: action.Primary}).Clicked() {
						ok = true
					}
				})
			}
			if closeButton(c, tk).Clicked() {
				*open = false
			}
		})
	})
	return ok
}

// The panel is built inside the backdrop, which appears in the same frame, so its Transition's
// Enter would never play: it scales in with a pose instead, and leaves inside the backdrop's exit.
func scaleIn(c *ui.Context, panel ui.Element) {
	tr := motion.Modal()
	entered := ui.Local(panel, "entered", func() bool { return false })
	pose := ui.Pose{}
	if !*entered {
		pose, *entered = ui.Pose{Scale: tr.Enter.Scale}, true
		c.Invalidate()
	}
	panel.AnimatePose(pose, ui.Animation{Duration: tr.Duration, Ease: tr.Ease})
}

func closeButton(c *ui.Context, tk *theme.Tokens) ui.Element {
	b := ui.ButtonBase(c).Label("Close").Absolute().Top(8).Right(12).Size(32, 32).Radius(12).
		Cursor(ui.CursorPointer).Transition(closeTransition)
	if b.Hovered() {
		b.Background(tk.FillSecondary).TextColor(tk.Text)
	} else {
		b.TextColor(tk.TextTertiary)
	}
	b.Children(func() { ui.Icon(c, lucide.X()).FontSize(18) })
	return b
}

func textOr(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}
