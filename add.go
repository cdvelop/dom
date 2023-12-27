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

	d.Log(d.setModuleActual(home_module))

	h.DevicePeripherals.KeyboardClientAdapter = d

	return d
}
