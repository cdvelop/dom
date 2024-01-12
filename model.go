package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type Dom struct {
	*model.MainHandler

	actualModule *model.Module // modulo cliqueado por usuario
	actualObject *model.Object //objeto en uso por usuario

	home_module string

	// cache vars
	elementJS js.Value
	stringVAR string //variable temporal
	objectID  string

	//keyboard
	keyboardFunc js.Func
	keyCode      int

	fnJsCall model.CallJsOptions
	err      string
}

type HtmlElement struct {
	Container js.Value
	Name      string //ej: div,li
	Id        string //ej: 234
	Class     string // ej: .css-class
	Content   string // ej: <h1>hello</h1>
}
