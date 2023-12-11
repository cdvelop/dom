package dom

import (
	"syscall/js"
)

// querySelector ej: "a[name='xxx']"
func (d Dom) ElementClicking(querySelector string) (err string) {
	const this = "ElementClicking error "
	element, err := query(querySelector)
	if err != "" {
		return this + err
	}

	result := element.Call("click")
	if result.Truthy() { //si retorna algo es por que ocurrió un error
		return this + result.String()
	}

	return ""
}

// ej: querySelector "meta[name='JsonBootTests']"
// get_content: "content"
func (d Dom) SelectContent(querySelector, get_content string) (content, err string) {
	const this = "SelectContent error "
	element, err := query(querySelector)
	if err != "" {
		return "", this + err
	}

	jsValue := element.Get(get_content)
	if !jsValue.Truthy() { //si retorna algo es por que ocurrió un error
		return "", this + "contenido: " + get_content + ", no encontrado con get"
	}

	return jsValue.String(), ""
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
