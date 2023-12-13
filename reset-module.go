package dom

import (
	"syscall/js"
)

func (d Dom) resetModule(v js.Value, p []js.Value) interface{} {
	const this = "resetModule error "
	if len(p) != 1 {
		return d.Log("error resetModule required 1 args")
	}

	module_name := p[0].Get("dataset").Get("module")
	if !module_name.Truthy() {
		d.UserMessage("error", "no se logro obtener nombre del modulo para volver a comenzar")
		return nil
	}

	d.Log("resetModule:", module_name.String())

	module, err := d.GetModuleByName(module_name.String())
	if err != "" {
		return this + err
	}

	module.ResetFrontendStateObjects()

	return nil

}
