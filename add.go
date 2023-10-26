package dom

import (
	"syscall/js"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/formclient"
	"github.com/cdvelop/logclient"
	"github.com/cdvelop/model"
	"github.com/cdvelop/timeclient"
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
		h:          h,
		cut:        &cutkey.Cut{},
		modules:    []*model.Module{},
		objects:    []*model.Object{},
		FormClient: nil,
		Logger:     logclient.Log{},
	}

	h.DomAdapter = &new
	h.TimeAdapter = timeclient.TimeCLient{}

	new.FormClient = formclient.Add(&new)

	return &new
}

func (d *Dom) AddModules(modules ...*model.Module) {

	objects := []*model.Object{}

	for _, m := range modules {
		objects = append(objects, m.Objects...)
	}

	d.modules = modules
	d.objects = objects

	d.h.CreateTablesInDB(objects, d)

	d.cut = cutkey.Add(objects...)
}
