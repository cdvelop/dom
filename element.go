package dom

func (e HtmlElement) Add() {
	if !e.Container.IsNull() && e.Name != "" {

		// Crear un nuevo elemento
		html_element := doc.Call("createElement", e.Name)

		if e.Id != "" {
			html_element.Set("id", e.Id)
		}

		if e.Class != "" {
			html_element.Set("className", e.Class)
		}

		if e.Content != "" {
			html_element.Set("innerHTML", e.Content)
		}

		// Agregar el nuevo elemento  al dom
		e.Container.Call("appendChild", html_element)

		// Log("Elemento " + e.Name + " agregado al dom")

	}

}
