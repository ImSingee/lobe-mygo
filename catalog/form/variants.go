package form

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/form"
	"github.com/lobehub/lobe-mygo/pkg/lucide"
)

type Prefs struct {
	Sounds, Badges, Digest, Compact bool
}

var (
	filledForm   = form.New(Prefs{Sounds: true, Badges: true})
	outlinedForm = form.New(Prefs{Digest: true})
)

func prefs(c *ui.Context, f *form.Form[Prefs], v form.Variant) {
	form.View(c, f, form.Props[Prefs]{Variant: v}, func() {
		form.Group(c, f, form.GroupProps{Title: "Notifications", Icon: lucide.Bell()}, func() {
			form.Field(c, f, form.FieldProps{Name: "sounds", Label: "Sounds", Desc: "Play a sound for new messages", Value: f.Values.Sounds}, func() {
				choice.Switch(c, &f.Values.Sounds, choice.SwitchProps{})
			})
			form.Field(c, f, form.FieldProps{Name: "badges", Label: "Badges", Value: f.Values.Badges}, func() {
				choice.Switch(c, &f.Values.Badges, choice.SwitchProps{})
			})
		})
		form.Group(c, f, form.GroupProps{Title: "Email", Icon: lucide.Mail(), DefaultCollapsed: true}, func() {
			form.Field(c, f, form.FieldProps{Name: "digest", Label: "Weekly digest", Value: f.Values.Digest}, func() {
				choice.Switch(c, &f.Values.Digest, choice.SwitchProps{})
			})
		})
	})
}

func Variants(c *ui.Context) {
	ui.Row(c).Gap(24).AlignItems(ui.Start).Children(func() {
		ui.Column(c.Key("filled")).Width(360).Children(func() { prefs(c, filledForm, form.Filled) })
		ui.Column(c.Key("outlined")).Width(360).Children(func() { prefs(c, outlinedForm, form.Outlined) })
	})
}
