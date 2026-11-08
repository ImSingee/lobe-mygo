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
	st := ui.Local(ui.Box(c).Absolute().Size(0, 0), "modal", func() modalState { return modalState{} })
	if *open {
		st.shown, st.closing = true, false
	} else if st.shown {
		st.closing = true
	}
	if !st.shown {
		return false
	}
	ok, gone := false, false
	ui.DialogBase(c, &st.shown, func(backdrop, panel ui.Element) {
		backdrop.Padding(32, 16)
		width := p.Width
		if width == 0 {
			width = 520
		}
		panel.Width(width).MaxWidthPercent(100).MaxHeightPercent(100).Radius(12).Clip().
			Background(tk.BgElevated).TextColor(tk.Text).FontSize(tk.FontSize)
		tk.Elevate(panel, theme.BoxShadow)
		gone = animateModal(tk, backdrop, panel, st)
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
	switch {
	case !st.shown:
		*open, st.shown, st.closing = false, true, true
		c.Invalidate()
	case gone:
		*st = modalState{}
		c.Invalidate()
	}
	return ok
}

// DialogBase stops building the dialog the frame it closes, and builds the panel inside the
// backdrop, which appears in the same frame, so neither the panel's Enter nor its Exit would play:
// the modal keeps the dialog built while it leaves and animates values it jumps to their start,
// which hold across the two builds of a frame a click makes, where a first pose would not. The
// backdrop animates its color, not its opacity, which the panel inside it would take too.
type modalState struct{ shown, closing, entered bool }

func animateModal(tk *theme.Tokens, backdrop, panel ui.Element, st *modalState) (gone bool) {
	m, b := motion.Modal(), motion.Backdrop()
	if !st.entered {
		st.entered = true
		panel.Jump("shown", 0)
		backdrop.Jump("dim", 0)
	}
	target, from, anim := float32(1), m.Enter.Scale, ui.Animation{Duration: m.Duration, Ease: m.Ease}
	if st.closing {
		target, from, anim = 0, m.Exit.Scale, ui.Animation{Duration: m.ExitDuration, Ease: m.ExitEase}
		backdrop.IgnoreInput(true)
	}
	shown := panel.AnimateTo("shown", target, anim)
	dim := backdrop.AnimateTo("dim", target, ui.Animation{Duration: b.Duration, Ease: b.Ease})
	scale := from + (1-from)*shown
	panel.Scale(scale, scale).Opacity(shown)
	backdrop.Background(tk.BgContainer.Alpha(0.6 * dim))
	return st.closing && shown <= 0 && dim <= 0
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
