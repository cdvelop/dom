package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type Dom struct {
	db      model.DataBaseAdapter
	theme   model.ThemeAdapter
	modules []*model.Module
	objects []*model.Object

	last_object *model.Object
	html_form   js.Value
	data_object map[string]string

	action_create bool
	action_update bool
	action_delete bool

	timeout_typing js.Value
}

type HtmlElement struct {
	Container js.Value
	Name      string //ej: div,li
	Id        string //ej: 234
	Class     string // ej: .css-class
	Content   string // ej: <h1>hello</h1>
}
