package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type Dom struct {
	*model.MainHandler

	moduleActual *model.Module
	objects      []*model.Object
	objectActual *model.Object

	home_module string

	objectJS   js.Value
	objectNAME string
	objectID   string

	err string
}

type HtmlElement struct {
	Container js.Value
	Name      string //ej: div,li
	Id        string //ej: 234
	Class     string // ej: .css-class
	Content   string // ej: <h1>hello</h1>
}
