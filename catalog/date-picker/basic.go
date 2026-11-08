package datepicker

import (
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

var (
	due      time.Time
	launch   = time.Date(2026, time.December, 1, 0, 0, 0, 0, time.Local)
	reminder time.Time
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(12).Width(260).Height(440).Children(func() {
		field.DatePicker(c, &due, field.DatePickerProps{AllowClear: true}).Label("Due date")
		field.DatePicker(c, &launch, field.DatePickerProps{Format: "Jan 2, 2006", Variant: field.InputFilled}).Label("Launch")
		field.DatePicker(c, &reminder, field.DatePickerProps{Size: theme.Small, Placeholder: "Remind me on"}).Label("Reminder")
	})
}
