package dom

import (
	"syscall/js"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/fetchclient"
	"github.com/cdvelop/formclient"
	"github.com/cdvelop/logclient"
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
		h:          h,
		cut:        &cutkey.Cut{},
		modules:    []*model.Module{},
		objects:    []*model.Object{},
		FormClient: nil,
		Logger:     logclient.Log{},
	}

	h.DomAdapter = &new
	h.ObjectAdapter = &new

	new.FormClient = formclient.Add(&new, h.DataBaseAdapter)

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

	http, err := fetchclient.Add(d.h.Logger, d.cut)
	if err != nil {
		d.Log(err)
	}

	// añadimos el controlador http
	d.h.FetchAdapter = http
}
