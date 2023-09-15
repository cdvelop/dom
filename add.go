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

func New(t model.Theme, db model.DataBaseAdapter) *Dom {

	return &Dom{
		db:      db,
		theme:   t,
		modules: nil,
		objects: nil,
	}
}

func (d *Dom) AddModules(modules ...*model.Module) {

	objects := []*model.Object{}

	for _, m := range modules {
		for _, o := range m.Objects {
			objects = append(objects, o)
		}
	}

	d.modules = modules
	d.objects = objects
}
