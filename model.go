package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type Dom struct {
	*model.MainHandler

	objects []*model.Object

	clickedModule *model.Module // modulo cliqueado por usuario
	clickedObject *model.Object //objeto cliqueado por usuario

	home_module string

	// cache vars
	objectJS  js.Value
	stringVAR string //variable temporal
	objectID  string

	err string
}

type HtmlElement struct {
	Container js.Value
	Name      string //ej: div,li
	Id        string //ej: 234
	Class     string // ej: .css-class
	Content   string // ej: <h1>hello</h1>
}
