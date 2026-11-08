package motion

import (
	"github.com/egoist/mygo/ui"
	lobe "github.com/lobehub/lobe-mygo"
	"github.com/lobehub/lobe-mygo/motion"
)

var shown = map[string]bool{}

func Presets(c *ui.Context) {
	tk := lobe.TokensOf(c)
	ui.Row(c).Gap(8).Wrap().Justify(ui.Center).Children(func() {
		for _, name := range []string{"Popup", "Tooltip", "Modal", "Drawer", "Toast", "Collapse"} {
			variant := lobe.Default
			if shown[name] {
				variant = lobe.Primary
			}
			if lobe.Button(c.Key(name), name, lobe.ButtonProps{Variant: variant, Size: lobe.Small}).Clicked() {
				shown[name] = !shown[name]
			}
		}
	})

	stage := ui.Box(c).AlignSelf(ui.Stretch).Height(260).Radius(tk.BorderRadiusLG).
		Background(tk.BgLayout).Border(1, tk.FillBorder()).Clip()
	stage.Children(func() {
		if shown["Collapse"] {
			ui.Column(c).Key("collapse").Absolute().Top(16).Left(16).Width(200).Padding(12).Gap(6).
				Radius(tk.BorderRadius).Background(tk.BgContainer).Transition(motion.Collapse()).
				Children(func() {
					ui.Text(c, "Collapse").FontWeight(600)
					ui.Text(c, "Height grows in while the content rises 6 DIPs.").FontSize(tk.FontSizeSM).TextColor(tk.TextSecondary)
				})
		}
		if shown["Popup"] {
			card(c, tk, "popup", "Popup").Absolute().Top(16).Left(420).Transition(motion.Popup())
		}
		if shown["Tooltip"] {
			ui.Text(c.Key("tooltip"), "Tooltip").Absolute().Top(110).Left(32).Padding(4, 8).
				Radius(tk.BorderRadiusSM).Background(tk.BgSpotlight).TextColor(ui.RGB(255, 255, 255)).
				FontSize(tk.FontSizeSM).Transition(motion.Tooltip())
		}
		if shown["Toast"] {
			card(c, tk, "toast", "Toast").Absolute().Bottom(16).Left(16).Transition(motion.Toast(motion.Bottom))
		}
		if shown["Modal"] {
			ui.Box(c).Key("backdrop").Absolute().Top(0).Left(0).Right(0).Bottom(0).
				Background(tk.BgMask).Transition(motion.Backdrop())
			card(c, tk, "modal", "Modal").Absolute().Top(80).Left(160).Width(220).Transition(motion.Modal())
		}
		if shown["Drawer"] {
			ui.Column(c).Key("drawer").Absolute().Top(0).Right(0).Bottom(0).Width(180).Padding(16).
				Background(tk.BgElevated).BorderWidth(0, 0, 0, 1).BorderColor(tk.FillBorder()).
				Transition(motion.Drawer(motion.Right)).
				Children(func() { ui.Text(c, "Drawer").FontWeight(600) })
		}
	})
}

func card(c *ui.Context, tk *lobe.Tokens, key, title string) ui.Element {
	return ui.Column(c).Key(key).Padding(12).Gap(4).Radius(tk.BorderRadius).
		Background(tk.BgElevated).Border(1, tk.FillBorder()).
		Shadow(0, 6, 16, 0, ui.RGBA(0, 0, 0, 0.12)).
		Children(func() {
			ui.Text(c, title).FontWeight(600)
			ui.Text(c, "lobe motion preset").FontSize(tk.FontSizeSM).TextColor(tk.TextSecondary)
		})
}
