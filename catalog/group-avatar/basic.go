package groupavatar

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

var members = []display.GroupAvatarItem{
	{Avatar: "Ada Lovelace"}, {Avatar: "🤖"}, {Avatar: "Grace Hopper"}, {Avatar: "🧠"},
	{Avatar: "Alan Turing"}, {Avatar: "✨"}, {Avatar: "Claude"},
}

func Basic(c *ui.Context) {
	ui.Column(c).Gap(20).Children(func() {
		ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
			for n := 1; n <= 4; n++ {
				display.GroupAvatar(c.Key(n), display.GroupAvatarProps{Size: 56, Avatars: members[:n]})
			}
			display.GroupAvatar(c.Key("auto"), display.GroupAvatarProps{Size: 56, Grid: display.GroupGridAuto, Avatars: members})
		})
		ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
			for i, corner := range []display.GroupCorner{display.GroupCornerSquare, display.GroupCornerCircle, display.GroupCornerSquircle, display.GroupCornerSmooth} {
				display.GroupAvatar(c.Key(i), display.GroupAvatarProps{Size: 56, Corner: corner, Avatars: members[:4]})
			}
			display.GroupAvatar(c.Key("round"), display.GroupAvatarProps{Size: 56, Corner: display.GroupCornerCircle,
				AvatarShape: display.AvatarCircle, Avatars: members[:4]})
		})
	})
}
