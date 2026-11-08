package home

import (
	_ "embed"
	"sync"

	"github.com/egoist/mygo/ui"
)

var (
	//go:embed logo.png
	logoPNG []byte
	//go:embed og-art.png
	ogArtPNG []byte
	//go:embed banner-art.png
	bannerArtPNG []byte
)

func decoded(b []byte) func() ui.ImageSource {
	return sync.OnceValue(func() ui.ImageSource {
		img, err := ui.DecodeImage(b)
		if err != nil {
			panic(err)
		}
		return img
	})
}

var (
	logo      = decoded(logoPNG)
	ogArt     = decoded(ogArtPNG)
	bannerArt = decoded(bannerArtPNG)
)

var (
	ogHeadline    = [2]string{"The UI Kit For", "Native Go Apps"}
	ogDescription = [2]string{"LobeHub's open-source Go component library for", "native desktop apps, drawn on the GPU."}
)

var bangs = []struct {
	color string
	tilt  float32
}{{"#589CBC", 7}, {"#E67183", 9}, {"#40A844", 4}, {"#B2913B", 11}, {"#7777AA", -8}}

func bangMarks(c *ui.Context, k float32) {
	ui.Row(c).Gap(7 * k).Children(func() {
		for i, b := range bangs {
			color := ui.Hex(b.color)
			ui.Column(c.Key(i)).Gap(5 * k).AlignItems(ui.Center).Rotate(b.tilt).Children(func() {
				ui.Box(c).Size(14*k, 46*k).Radius(2 * k).Background(color)
				ui.Box(c).Size(14*k, 13*k).Radius(2 * k).Background(color)
			})
		}
	})
}

// OG is the site's social preview, after lobe-ui's: the name and promise on the left over an
// illustration of the LobeHub mascot juggling MyGo's five marks.
func OG(c *ui.Context) {
	ui.Box(c).Fill().Background(ui.Hex("#050505")).Clip().Children(func() {
		ui.Image(c, ogArt()).Absolute().Left(0).Top(0).Size(1280, 640).Label("art")
		ogText(c)
	})
}

// Banner is the README's header, after lobe-ui's: the social card's text over a wider cut of the
// art, on a rounded card whose corners stay transparent.
func Banner(c *ui.Context) {
	ui.Box(c).Size(1474, 602).Radius(24).Background(ui.Hex("#050505")).Clip().Children(func() {
		ui.Image(c, bannerArt()).Absolute().Left(0).Top(0).Size(1474, 602).Label("banner-art")
		ogText(c)
		ui.Box(c).Absolute().Left(0).Top(0).Right(0).Bottom(0).Radius(24).Border(1, ui.RGBA(255, 255, 255, 0.12)).IgnoreInput(true)
	})
}

func ogText(c *ui.Context) {
	white, grey := ui.Hex("#FFFFFF"), ui.Hex("#9B9BA1")
	ui.Column(c).Absolute().Left(64).Top(104).Children(func() {
		ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
			ui.Image(c, logo()).Size(56, 56)
			ui.Text(c, "/").FontSize(40).TextColor(ui.Hex("#5A5A60"))
			bangMarks(c, 0.62)
			ui.Text(c, "lobe-mygo").Font("Menlo").FontSize(34).FontWeight(700).TextColor(white)
		})
	})
	ui.Column(c).Absolute().Left(64).Top(222).Children(func() {
		for i, line := range ogHeadline {
			ui.Text(c.Key(i), line).Font("Helvetica Neue").FontSize(76).FontWeight(700).LetterSpacing(-1.5).LineHeight(1.12).TextColor(white)
		}
	})
	ui.Column(c).Absolute().Left(64).Top(486).Children(func() {
		for i, line := range ogDescription {
			ui.Text(c.Key(i), line).FontSize(23).LineHeight(1.5).TextColor(grey)
		}
	})
}
