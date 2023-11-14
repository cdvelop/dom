package dom

import (
	"syscall/js"

	"github.com/cdvelop/formclient"
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

	new := Dom{
		ThemeAdapter:    h,
		Logger:          h,
		DataBaseAdapter: h,
		DataConverter:   h,
		FormClient:      nil,
		ObjectsHandler:  h,
		ModuleHandler:   h,
		FetchAdapter:    h,
	}

	h.DomAdapter = new
	new.FormClient = formclient.Add(&new, h.DataBaseAdapter)

	return &new
}
