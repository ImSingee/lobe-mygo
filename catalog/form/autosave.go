package form

import (
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/choice"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/form"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

type Agent struct {
	Name   string
	Stream bool
}

var (
	agentForm = form.New(Agent{Name: "Translator", Stream: true})
	saved     = "Not changed yet"
	settings  = form.New(Agent{Name: "Writer"})
)

func Autosave(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Row(c).Gap(32).AlignItems(ui.Start).Children(func() {
		ui.Column(c.Key("autosave")).Width(340).Gap(8).Children(func() {
			f := agentForm
			form.View(c, f, form.Props[Agent]{ChangeDebounce: 500 * time.Millisecond, OnValuesChange: func(a Agent) {
				saved = "Saved “" + a.Name + "” at " + time.Now().Format("15:04:05")
			}}, func() {
				form.Group(c, f, form.GroupProps{Title: "Auto-save"}, func() {
					form.Field(c, f, form.FieldProps{Name: "name", Label: "Name", Value: f.Values.Name}, func() {
						field.Input(c, &f.Values.Name, field.InputProps{}).Width(160)
					})
					form.Field(c, f, form.FieldProps{Name: "stream", Label: "Stream replies", Value: f.Values.Stream}, func() {
						choice.Switch(c, &f.Values.Stream, choice.SwitchProps{})
					})
				})
			})
			ui.Text(c, saved).FontSize(12).TextColor(tk.TextSecondary)
		})
		ui.Column(c.Key("float")).Width(340).Children(func() {
			f := settings
			form.View(c, f, form.Props[Agent]{OnSubmit: func(Agent) error { return nil }}, func() {
				form.Group(c, f, form.GroupProps{Title: "Floating footer"}, func() {
					form.Field(c, f, form.FieldProps{Name: "name", Label: "Name", Desc: "Edit to show the footer", Value: f.Values.Name}, func() {
						field.Input(c, &f.Values.Name, field.InputProps{}).Width(160)
					})
				})
				form.SubmitFooter(c, f, form.SubmitFooterProps{Float: true, Texts: form.FooterTexts{Submit: "Save"}})
			})
		})
	})
}
