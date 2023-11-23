package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) GetHtmlContainer(o *model.Object) (ctn *js.Value, err string) {

	container := doc.Call("querySelector", d.QuerySelectorObject(o.ModuleName, o.ObjectName))
	if container.Truthy() {
		return &container, ""
	}

	return nil, "GetHtmlContainer error no se logro obtener contenedor objeto: " + o.ObjectName
}

func (d Dom) GetHtmlModule(module_name string) (out any, err string) {

	module_html := body.Call("querySelector", "div#"+module_name)
	if !module_html.Truthy() {
		return nil, "GetHtmlModule error. modulo html " + module_name + " no encontrado"
	}

	return module_html, ""

}
