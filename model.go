package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type Dom struct {
	db      model.DataBaseAdapter
	theme   model.Theme
	modules []*model.Module
	header  model.FrontendHeaderHandler
	// object_id + object
	objects map[string]*model.Object

	last_object *model.Object
	form        js.Value
	data_object map[string]string

	timeout_typing js.Value
}

type HtmlElement struct {
	Container js.Value
	Name      string //ej: div,li
	Id        string //ej: 234
	Class     string // ej: .css-class
	Content   string // ej: <h1>hello</h1>
}
