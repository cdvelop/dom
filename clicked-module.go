package dom

import "syscall/js"

func (d Dom) moduleClickedUI(t js.Value, module []js.Value) interface{} {
	const e = " func moduleClickedUI error"
	if len(module) != 1 {
		return d.Log("se requiere un argumento" + e)
	}

	d.Log("CLICK MODULO:", module[0].String(), e)

	d.err = d.SetActualModule(module[0].String())
	if d.err != "" {
		return d.Log(d.err + e)
	}

	return nil

}
