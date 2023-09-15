package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func getHtmlModule(module_name string) (*js.Value, error) {

	module_html := body.Call("querySelector", "#"+module_name)
	if !module_html.Truthy() {
		return nil, model.Error("error modulo html", module_name, "no encontrado")
	}

	return &module_html, nil

}
