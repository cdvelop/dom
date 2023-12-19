package dom

import "syscall/js"

func (d *Dom) moduleClickedUI(t js.Value, module []js.Value) interface{} {
	const e = ". moduleClickedUI"
	if len(module) != 1 {
		return d.Log("se requiere un argumento" + e)
	}

	d.Log(d.setModuleActual(module[0].String()))

	return nil

}

func (d *Dom) setModuleActual(module_name string) (err string) {
	const e = ". setModuleActual"

	d.Log("CLICK MODULO:", module_name, e)

	if d.clickedModule != nil && d.clickedModule.ModuleName == module_name {
		return "" // nada que hacer
	}

	d.clickedModule, d.err = d.GetModuleByName(module_name)
	if d.err != "" {
		return d.err + e
	}

	return ""
}
