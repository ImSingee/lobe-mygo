package form

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/form"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

type Profile struct {
	Name, Bio string
	Beta      bool
}

var profileForm = form.New(Profile{Name: "Ada"})

func Basic(c *ui.Context) {
	f := profileForm
	ui.Column(c).Width(600).Children(func() {
		form.View(c, f, form.Props[Profile]{ItemMinWidth: 240, OnSubmit: func(p Profile) error {
			feedback.ShowToast(c, feedback.ToastProps{Title: "Saved " + p.Name, Type: feedback.ToastSuccess})
			return nil
		}}, func() {
			form.Group(c, f, form.GroupProps{Title: "Profile", Icon: lucide.Settings()}, func() {
				form.Field(c, f, form.FieldProps{Name: "name", Label: "Name", Desc: "Shown to other members", Required: true, Value: f.Values.Name}, func() {
					field.Input(c, &f.Values.Name, field.InputProps{Invalid: f.Error("name") != ""}).Width(240)
				})
				form.Field(c, f, form.FieldProps{Name: "bio", Label: "Bio", Value: f.Values.Bio}, func() {
					field.TextArea(c, &f.Values.Bio, field.TextAreaProps{Placeholder: "Introduce yourself", MinRows: 2, MaxRows: 4}).Width(320)
				})
			})
			form.Group(c, f, form.GroupProps{Title: "Labs", Icon: lucide.Sparkles()}, func() {
				form.Field(c, f, form.FieldProps{Name: "beta", Label: "Beta features", Tag: "beta", Desc: "Enable experimental features", Value: f.Values.Beta}, func() {
					choice.Switch(c, &f.Values.Beta, choice.SwitchProps{})
				})
			})
			form.SubmitFooter(c, f, form.SubmitFooterProps{})
		})
		feedback.ToastHost(c, feedback.ToastHostProps{})
	})
}
