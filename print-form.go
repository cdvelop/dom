package dom

import "syscall/js"

func (d *Dom) printForm(this js.Value, p []js.Value) interface{} {
	const e = ". printForm"
	if len(p) != 1 {
		return d.Log("required 1 args" + e)
	}

	d.objectJS = p[0].Get("dataset").Get("object_name")
	if !d.objectJS.Truthy() {
		d.UserMessage("error no se logro obtener nombre del formulario para imprimir")
		return nil
	}

	d.err = d.SetActualObject(d.objectJS.String())
	if d.err != "" {
		d.UserMessage(d.err)
		return nil
	}

	if d.objectActual.PrinterHandler == nil {
		d.UserMessage("err", d.objectActual.Title, "no cuenta con controlador para imprimir")
		return nil
	}

	d.objectActual.PrintFormObject()

	return nil

}
