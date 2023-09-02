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

func New(t model.Theme, a model.FrontendAuthHandler, modules ...*model.Module) *Dom {

	var auth model.FrontendAuthHandler
	if a == nil {
		auth = model.DefaultAuthHandler{}
	}

	objects := map[string]*model.Object{}

	for _, m := range modules {
		for _, o := range m.Objects {
			objects[o.ID()] = o
		}
	}

	return &Dom{
		db:      indexdb.Add(),
		theme:   t,
		modules: modules,
		auth:    auth,
		objects: objects,
	}
}
