package togglegroup

import (
	"strings"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

var (
	align  = []string{"left"}
	format = []string{"bold"}
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Children(func() {
		choice.ToggleGroup(c, &align, choice.ToggleGroupProps{Items: []choice.ToggleItem{
			{Value: "left", Icon: lucide.AlignLeft()},
			{Value: "center", Icon: lucide.AlignCenter()},
			{Value: "right", Icon: lucide.AlignRight()},
		}})
		choice.ToggleGroup(c, &format, choice.ToggleGroupProps{Multiple: true, Items: []choice.ToggleItem{
			{Value: "bold", Icon: lucide.Bold(), Label: "Bold"},
			{Value: "italic", Icon: lucide.Italic(), Label: "Italic"},
			{Value: "underline", Icon: lucide.Underline(), Label: "Underline"},
		}})
		ui.Text(c, "Align: "+strings.Join(align, ", ")+" · Format: "+strings.Join(format, ", ")).FontSize(12)
	})
}
