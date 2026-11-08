package form

import (
	"context"
	"strings"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/field"
	"github.com/lobehub/lobe-mygo/pkg/form"
)

type Account struct {
	Email, Slug, Password, Confirm string
}

var accountForm = form.New(Account{})

func checkSlug(ctx context.Context, v any) string {
	select {
	case <-ctx.Done():
		return ""
	case <-time.After(600 * time.Millisecond):
	}
	switch strings.ToLower(v.(string)) {
	case "admin", "root", "lobe":
		return "This slug is taken"
	}
	return ""
}

func Validation(c *ui.Context) {
	f := accountForm
	invalid := func(name string) bool { return f.Error(name) != "" }
	ui.Column(c).Width(480).Children(func() {
		form.View(c, f, form.Props[Account]{Layout: form.Vertical}, func() {
			form.Group(c, f, form.GroupProps{Title: "Create account"}, func() {
				form.Field(c, f, form.FieldProps{Name: "email", Label: "Email", Required: true, Value: f.Values.Email,
					Validate: func() string {
						if !strings.Contains(f.Values.Email, "@") {
							return "Enter an email address"
						}
						return ""
					}}, func() {
					field.Input(c, &f.Values.Email, field.InputProps{Placeholder: "you@example.com", Invalid: invalid("email")})
				})
				form.Field(c, f, form.FieldProps{Name: "slug", Label: "Slug", Desc: "Checked against taken names as you type", Value: f.Values.Slug,
					ValidateOn: form.ValidateChange, ValidateAsync: checkSlug, ValidateDebounce: 300 * time.Millisecond}, func() {
					suffix := ""
					if f.Validating("slug") {
						suffix = "Checking…"
					}
					field.Input(c, &f.Values.Slug, field.InputProps{Placeholder: `try "admin"`, Invalid: invalid("slug")})
					if suffix != "" {
						ui.Text(c, suffix).FontSize(12)
					}
				})
				form.Field(c, f, form.FieldProps{Name: "password", Label: "Password", Required: true, Value: f.Values.Password,
					Validate: func() string {
						if len(f.Values.Password) < 8 {
							return "At least 8 characters"
						}
						return ""
					}}, func() {
					field.InputPassword(c, &f.Values.Password, field.InputProps{Invalid: invalid("password")})
				})
				form.Field(c, f, form.FieldProps{Name: "confirm", Label: "Confirm password", Required: true, Value: f.Values.Confirm, Deps: []string{"password"},
					Validate: func() string {
						if f.Values.Confirm != f.Values.Password {
							return "The passwords differ"
						}
						return ""
					}}, func() {
					field.InputPassword(c, &f.Values.Confirm, field.InputProps{Invalid: invalid("confirm")})
				})
			})
			form.SubmitFooter(c, f, form.SubmitFooterProps{Texts: form.FooterTexts{Submit: "Create account"}})
		})
	})
}
