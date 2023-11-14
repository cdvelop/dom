package dom

import "syscall/js"

func (d Dom) resetModule(this js.Value, p []js.Value) interface{} {

	if len(p) != 1 {
		return d.Log("error resetModule required 1 args")
	}

	module_name := p[0].Get("dataset").Get("module")
	if !module_name.Truthy() {
		d.UserMessage("error", "no se logro obtener nombre del modulo para volver a comenzar")
		return nil
	}

	d.Log(" reset module_name:", module_name)

	module, err := d.GetModuleByName(module_name.String())
	if err != nil {
		return err
	}

	for _, o := range module.Objects {
		if o.ViewReset != nil {
			o.ViewReset.ResetView()
		}
	}

	return nil

}
