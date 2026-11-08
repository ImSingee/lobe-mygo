package toast

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/action"
	"github.com/lobehub/lobe-mygo/pkg/feedback"
)

func Basic(c *ui.Context) {
	kinds := []struct {
		label string
		toast feedback.ToastProps
	}{
		{"Success", feedback.ToastProps{Type: feedback.ToastSuccess, Title: "Changes saved", Description: "Everyone with access sees them now."}},
		{"Info", feedback.ToastProps{Type: feedback.ToastInfo, Title: "Syncing", Description: "Your files are on their way."}},
		{"Warning", feedback.ToastProps{Type: feedback.ToastWarning, Title: "Storage almost full"}},
		{"Error", feedback.ToastProps{Type: feedback.ToastError, Title: "Upload failed", Description: "Check your connection and try again."}},
	}
	ui.Row(c).Gap(8).Children(func() {
		for _, k := range kinds {
			if action.Button(c.Key(k.label), k.label, action.ButtonProps{}).Clicked() {
				feedback.ShowToast(c, k.toast)
			}
		}
	})
	feedback.ToastHost(c, feedback.ToastHostProps{})
}
