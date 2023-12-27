package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

// querySelector ej: "a[name='xxx']"
func (d Dom) ElementClicking(querySelector string) (err string) {
	return d.elementCall("ElementClicking", querySelector, "click")
}

func (d Dom) ElementFocus(querySelector string) (err string) {
	return d.elementCall("ElementFocus", querySelector, "focus")
}

// call ej: focus,click
func (d Dom) elementCall(e, querySelector, call string) (err string) {

	d.elementJS, err = query(querySelector)
	if err != "" {
		return e + " " + err
	}

	result := d.elementJS.Call(call)
	if result.Truthy() { //si retorna algo es por que ocurrió un error
		return e + " " + result.String()
	}

	return ""
}

// ej: querySelector "meta[name='JsonBootTests']"
// get_content: "content"
// set_after true = element.Set("content", "")
func (d Dom) SelectContent(o model.SelectDomOptions) (out any, err string) {
	const t = "SelectContent error "
	element, err := query(o.QuerySelector)
	if err != "" {
		return "", t + err
	}

	var jsValue js.Value
	if o.GetContent != "" {
		jsValue = element.Get(o.GetContent)
		if !jsValue.Truthy() { //si retorna algo es por que ocurrió un error
			return "", t + "contenido: " + o.GetContent + ", no encontrado con get"
		}

	} else { // si este vació o.GetContent quiere decir que solo necesita el elemento
		jsValue = element
	}

	if o.StringReturn {
		out = jsValue.String()
	} else {
		out = jsValue
	}

	if o.SetAfter {
		jsValue.Set(o.GetContent, "")
	}

	return
}

func query(selector string) (element js.Value, err string) {
	element = doc.Call("querySelector", selector)
	// d.Log("ELEMENTO CLICK", element)
	if !element.Truthy() {
		return js.Value{}, "query no se encontró elemento con la consulta " + selector
	}
	return element, ""
}

func (d Dom) GetHtmlModule(module_name string) (out any, err string) {

	module_html := body.Call("querySelector", "div#"+module_name)
	if !module_html.Truthy() {
		return nil, "GetHtmlModule error. modulo html " + module_name + " no encontrado"
	}

	return module_html, ""

}
