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

func New(h *model.Handlers) (*Dom, error) {

	dom := Dom{
		ThemeAdapter:    h,
		Logger:          h,
		DataBaseAdapter: h,
		DataConverter:   h,
		ObjectsHandler:  h,
		ModuleHandler:   h,
		FetchAdapter:    h,
	}

	h.DomAdapter = dom
	h.MessageAdapter = dom

	err := h.CheckInterfaces("dom config", dom)
	if err != nil {
		return nil, err
	}

	return &dom, nil
}
