package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type Dom struct {
	theme   model.Theme
	modules []*model.Module
	header  model.FrontendHeaderHandler
	// object_id + object
	objects map[string]*model.Object

	last_object *model.Object
}

type HtmlElement struct {
	Container js.Value
	Name      string //ej: div,li
	Id        string //ej: 234
	Class     string // ej: .css-class
	Content   string // ej: <h1>hello</h1>
}
