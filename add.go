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

func New(h *model.Handlers) *Dom {

	dom := Dom{
		Handlers: h,
	}

	h.DomAdapter = dom
	h.MessageAdapter = dom
	h.HtmlAdapter = dom

	return &dom
}
