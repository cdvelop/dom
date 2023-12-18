package dom

import "syscall/js"

func (d Dom) printForm(this js.Value, p []js.Value) interface{} {

	if len(p) != 1 {
		return d.Log("error printForm required 1 args")
	}

	object_name := p[0].Get("dataset").Get("object_name")
	if !object_name.Truthy() {
		d.UserMessage("error", "no se logro obtener nombre del formulario para imprimir")
		return nil
	}

	o, err := d.MainHandlerGetObjectByName(object_name.String())
	if err != "" {
		d.UserMessage(err)
		return nil
	}

	if o.PrinterHandler == nil {
		d.UserMessage("err", o.Title, "no cuenta con controlador para imprimir")
		return nil
	}

	o.PrintFormObject()

	return nil

}
