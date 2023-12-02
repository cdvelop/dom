package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

var (
	window, doc, body js.Value
)

func init() {
	window = js.Global()
	doc = window.Get("document")
	body = doc.Get("body")
}

// home_module ej:session
func New(h *model.Handlers, home_module string) *Dom {

	dom := Dom{
		Handlers:    h,
		home_module: home_module,
	}

	h.DomAdapter = dom
	h.MessageAdapter = dom
	h.HtmlAdapter = dom

	return &dom
}
