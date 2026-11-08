package motion

import (
	"time"

	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
	"github.com/lobehub/lobe-mygo/motion"
)

var played bool

func Springs(c *ui.Context) {
	tk := lobe.TokensOf(c)
	springs := []struct {
		name string
		anim ui.Animation
	}{
		{"Smooth", motion.Smooth(600 * time.Millisecond)},
		{"Snappy", motion.Snappy(600 * time.Millisecond)},
		{"Bouncy", motion.Bouncy(600 * time.Millisecond)},
		{"SwitchThumb", motion.SwitchThumb},
	}
	ui.Column(c).Gap(10).Width(420).Children(func() {
		for _, s := range springs {
			ui.Row(c.Key(s.name)).Gap(12).Children(func() {
				ui.Text(c, s.name).Width(90).FontSize(tk.FontSizeSM).TextColor(tk.TextSecondary)
				track := ui.Box(c).Grow(1).Height(20).Radius(10).Background(tk.FillTertiary)
				track.Children(func() {
					dot := ui.Box(c).Size(20, 20).Radius(10).Background(tk.Primary)
					target := float32(0)
					if played {
						target = 290
					}
					dot.Translate(dot.AnimateTo("x", target, s.anim), 0)
				})
			})
		}
	})
	if lobe.Button(c, "Play", lobe.ButtonProps{Variant: lobe.Primary}).Clicked() {
		played = !played
	}
}
