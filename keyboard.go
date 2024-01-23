package dom

import "syscall/js"

func (d *Dom) keyboardHandler(this js.Value, p []js.Value) interface{} {

	const e = "keyboardHandler error"

	// d.Log(e, "recuperar nombre objeto:", p[0])
	d.err = d.setActualObject(p[0].Get("target").Get("form").Get("name").String())
	if d.err != "" {
		d.Log(e + d.err)
		return nil
	}

	// Obtiene el código de la tecla presionada
	d.keyCode = p[0].Get("keyCode").Int()

	switch d.keyCode {
	case 13:
		// Previene la acción por defecto de la tecla Enter
		p[0].Call("preventDefault")

		d.KeyboardClientDisable(true)

		if d.actualObject.KeyboardHandlerObject.KeyEnterAdapter != nil {
			d.actualObject.KeyboardHandlerObject.KeyEnter()
		}

		// d.Log("TECLA ENTER PRESIONADA  (código 13)")
	default:
		d.Log("info KEYBOARD CODE:", d.keyCode)

	}

	return nil
}

func (d *Dom) KeyboardClientDisable(disable bool) {

	if disable {

		d.Log("info KEYBOARD OFF")
		// Elimina el evento d.keyboardCallback
		window.Get("document").Call("removeEventListener", "keydown", d.keyboardFunc)
		d.keyboardFunc.Release()

	} else {
		d.keyboardFunc = js.FuncOf(d.keyboardHandler)

		d.Log("info KEYBOARD ON")
		window.Get("document").Call("addEventListener", "keydown", d.keyboardFunc)

	}

}
