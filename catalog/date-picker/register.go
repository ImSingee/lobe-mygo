package datepicker

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("date-picker/basic", Basic)
	registry.Add("date-picker/calendar", CalendarDemo)
	registry.Add("date-picker/disabled", Disabled)
}
