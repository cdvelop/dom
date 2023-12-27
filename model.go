package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type Dom struct {
	*model.MainHandler

	clickedModule *model.Module // modulo cliqueado por usuario
	clickedObject *model.Object //objeto cliqueado por usuario

	home_module string

	// cache vars
	elementJS js.Value
	stringVAR string //variable temporal
	objectID  string

	//keyboard
	keyboardFunc js.Func
	keyCode      int

	err string
}

type HtmlElement struct {
	Container js.Value
	Name      string //ej: div,li
	Id        string //ej: 234
	Class     string // ej: .css-class
	Content   string // ej: <h1>hello</h1>
}
