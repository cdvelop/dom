package dom

import (
	"syscall/js"
)

func (d *Dom) userTyping(this js.Value, p []js.Value) interface{} {

	if d.timeout_typing.Truthy() {
		// Si hay un temporizador en curso, lo cancelamos
		js.Global().Call("clearTimeout", d.timeout_typing)
	}

	// Configuramos un nuevo temporizador para 500 milisegundos
	d.timeout_typing = js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, null []js.Value) interface{} {

		// Log("ejecutando acción después de 500 milisegundos")

		err := d.currentObject(p)
		if err != nil {
			Log(err.Error())
			return nil
		}

		err = d.validateForm(p[0])
		if err != nil {
			Log(err.Error())
			return nil
		}

		return nil
	}), 500)

	return nil

}
