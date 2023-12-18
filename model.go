package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type Dom struct {
	*model.MainHandler
	home_module string

	object_name string
	object_id   string

	err string
}

type HtmlElement struct {
	Container js.Value
	Name      string //ej: div,li
	Id        string //ej: 234
	Class     string // ej: .css-class
	Content   string // ej: <h1>hello</h1>
}
