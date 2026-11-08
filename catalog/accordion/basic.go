package accordion

import (
	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/display"
)

var (
	faqOpen = map[string]bool{}
	faq     = []struct{ key, q, a string }{
		{"what", "What is Lobe UI?", "Lobe UI is an open-source component library for AI apps."},
		{"go", "Does it run in Go?", "This port draws every component with mygo, natively and in the browser through WebAssembly."},
		{"theme", "Can I change the theme?", "The components read the Lobe UI tokens, light and dark, from the theme package."},
	}
)

func Basic(c *ui.Context) {
	ui.Column(c).Width(480).Height(200).Children(func() {
		items := make([]display.AccordionItem, len(faq))
		for i, f := range faq {
			items[i] = display.AccordionItem{Key: f.key, Title: f.q}
		}
		display.Accordion(c, faqOpen, display.AccordionProps{Items: items, Single: true}, func(key string) {
			for _, f := range faq {
				if f.key == key {
					ui.Text(c, f.a)
				}
			}
		})
	})
}
