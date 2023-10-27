package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) GetHtmlContainer(o *model.Object) (*js.Value, error) {

	container := doc.Call("querySelector", d.h.QuerySelectorObject(o.ModuleName, o.Name))
	if container.Truthy() {
		return &container, nil
	}

	return nil, model.Error("error no se logro obtener contenedor objeto:", o.Name)
}

func (Dom) GetHtmlModule(module_name string) (*js.Value, error) {

	module_html := body.Call("querySelector", "#"+module_name)
	if !module_html.Truthy() {
		return nil, model.Error("error modulo html", module_name, "no encontrado")
	}

	return &module_html, nil

}
