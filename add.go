package dom

import (
	"syscall/js"

	"github.com/cdvelop/indexdb"
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

func New(h *model.ModuleHandlers) *Dom {

	new := Dom{
		db:             h.DBA,
		theme:          h.Theme,
		modules:        []*model.Module{},
		objects:        []*model.Object{},
		last_object:    &model.Object{},
		html_form:      js.Value{},
		data_object:    map[string]string{},
		action_create:  false,
		action_update:  false,
		action_delete:  false,
		timeout_typing: js.Value{},
	}

	h.DOM = new

	return &new
}

func (d *Dom) AddModules(modules ...*model.Module) {

	objects := []*model.Object{}

	for _, m := range modules {
		objects = append(objects, m.Objects...)
	}

	d.modules = modules
	d.objects = objects

	d.db = indexdb.Add(objects)
}
