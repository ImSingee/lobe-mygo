package datepicker

import (
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
)

var picked = time.Date(2026, time.November, 8, 0, 0, 0, 0, time.Local)

func CalendarDemo(c *ui.Context) {
	ui.Column(c).Gap(8).Children(func() {
		field.Calendar(c, &picked, field.CalendarProps{})
		ui.Text(c, picked.Format("Monday, January 2, 2006")).FontSize(12)
	})
}
