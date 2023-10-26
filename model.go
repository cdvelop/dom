package dom

import (
	"syscall/js"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/formclient"
	"github.com/cdvelop/model"
)

type Dom struct {
	h *model.Handlers

	cut *cutkey.Cut

	modules []*model.Module
	objects []*model.Object

	*formclient.FormClient

	model.Logger
}

type HtmlElement struct {
	Container js.Value
	Name      string //ej: div,li
	Id        string //ej: 234
	Class     string // ej: .css-class
	Content   string // ej: <h1>hello</h1>
}
