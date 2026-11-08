package datepicker

import (
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
)

var meeting time.Time

func Disabled(c *ui.Context) {
	ui.Column(c).Width(260).Height(440).Children(func() {
		field.DatePicker(c, &meeting, field.DatePickerProps{
			Placeholder: "Pick a weekday",
			WeekStart:   time.Monday,
			DisabledDate: func(d time.Time) bool {
				return d.Weekday() == time.Saturday || d.Weekday() == time.Sunday
			},
		}).Label("Meeting")
	})
}
