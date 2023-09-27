package dom

import (
	"syscall/js"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

type Dom struct {
	h *model.ModuleHandlers

	cut *cutkey.Cut

	modules []*model.Module
	objects []*model.Object

	last_object *model.Object
	html_form   js.Value          //form
	data_object map[string]string //form

	action_create bool //form
	action_update bool //form
	action_delete bool //form

	timeout_typing js.Value //form
}

type HtmlElement struct {
	Container js.Value
	Name      string //ej: div,li
	Id        string //ej: 234
	Class     string // ej: .css-class
	Content   string // ej: <h1>hello</h1>
}
