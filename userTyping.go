package dom

import (
	"syscall/js"
)

func (d *Dom) userTyping(this js.Value, p []js.Value) interface{} {

	err := d.currentObject(p)
	if err != nil {
		Log(err.Error())
		return nil
	}

	err = d.validateForm(p[0], p[1])
	if err != nil {
		Log(err.Error())
		return nil
	}

	// js.Global().Get("console").Call("log", "VALIDANDO CAMPO:", fieldValue)

	// fmt.Println("OBJETO ENCONTRADO: ", object.Name)
	// hex.EncodeToString([]byte("OBJETO ENCONTRADO: " + object.Name))

	return nil
}
