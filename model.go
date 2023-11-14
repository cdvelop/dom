package dom

import (
	"syscall/js"

	"github.com/cdvelop/formclient"
	"github.com/cdvelop/model"
)

type Dom struct {
	model.ThemeAdapter
	model.Logger
	model.DataBaseAdapter
	model.DataConverter
	model.ObjectsHandler
	model.ModuleHandler
	model.FetchAdapter

	*formclient.FormClient
}

type HtmlElement struct {
	Container js.Value
	Name      string //ej: div,li
	Id        string //ej: 234
	Class     string // ej: .css-class
	Content   string // ej: <h1>hello</h1>
}
