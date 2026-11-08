package tokens

import (
	"fmt"
	"github.com/lobehub/lobe-mygo/pkg/theme"
	"strconv"
	"strings"

	"github.com/egoist/mygo/ui"
)

var groups = []string{"Text", "Bg", "Fill", "Border", "Primary", "Error", "Warning", "Success", "Info"}

func Swatches(c *ui.Context) {
	tk := theme.TokensOf(c)
	names := theme.ColorTokenNames()
	ui.Scroll(c).AlignSelf(ui.Stretch).Grow(1).Children(func() {
		ui.Column(c).Gap(24).Padding(4).Children(func() {
			for _, group := range groups {
				ui.Column(c).Gap(12).Children(func() {
					ui.Text(c, group).FontSize(tk.FontSizeLG).FontWeight(600)
					ui.Row(c).Wrap().Gap(12).Children(func() {
						for _, name := range names {
							if strings.HasPrefix(name, group) {
								color, _ := tk.Color(name)
								swatch(c, tk, name, color)
							}
						}
					})
				})
			}
		})
	})
}

func swatch(c *ui.Context, tk *theme.Tokens, name string, color ui.Color) {
	ui.Row(c.Key(name)).Width(216).Gap(10).Children(func() {
		ui.Box(c).Size(36, 36).Radius(tk.BorderRadiusSM).Background(color).Border(1, tk.FillBorder())
		ui.Column(c).Gap(2).Children(func() {
			ui.Text(c, name).FontSize(tk.FontSizeSM).FontWeight(500)
			ui.Text(c, cssValue(color)).FontSize(tk.FontSizeSM).TextColor(tk.TextTertiary)
		})
	})
}

func cssValue(c ui.Color) string {
	if c.A == 255 {
		return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
	}
	alpha := strconv.FormatFloat(float64(c.A)/255, 'f', 2, 64)
	return fmt.Sprintf("rgba(%d, %d, %d, %s)", c.R, c.G, c.B, alpha)
}
