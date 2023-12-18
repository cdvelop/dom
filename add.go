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
func New(h *model.MainHandler, home_module string) *Dom {

	d := &Dom{
		MainHandler: h,
		home_module: home_module,
	}

	h.DomAdapter = d
	h.MessageAdapter = d
	h.HtmlAdapter = d
	h.ObjectHandlerAdapter = d

	d.err = d.SetActualModule(home_module)
	if d.err != "" {
		d.Log("dom new error:", d.err)
	}

	return d
}
