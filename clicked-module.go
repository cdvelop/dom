package dom

import (
	"syscall/js"

	"github.com/cdvelop/strings"
)

func (d *Dom) moduleClickedUI(t js.Value, evt []js.Value) interface{} {
	const e = ". moduleClickedUI"
	if len(evt) != 1 {
		return d.Log("se requiere un argumento" + e)
	}

	d.elementJS = evt[0].Get("target")
	d.stringVAR = d.elementJS.Get("tagName").String()

	switch strings.ToLowerCase(d.stringVAR) { //tagName
	case "svg", "span":
		d.elementJS = d.elementJS.Get("parentNode")
	case "path", "use":
		d.elementJS = d.elementJS.Get("parentNode").Get("parentNode")
	}

	// d.Log("TARGET MODULE:", d.elementJS)

	d.stringVAR = d.elementJS.Get("name").String() //nombre modulo

	d.err = d.setModuleActual(d.stringVAR)
	if d.err != "" {
		d.Log(d.err + e)
	}
	// d.Log("MODULO ACTUAL:", d.actualModule.ModuleName)

	if d.actualModule.FrontendModuleHandlers.ClickedModuleEventAdapter != nil {
		d.actualModule.FrontendModuleHandlers.ClickedModuleEvent()
	}

	d.menuRouter(d.elementJS)

	// _, d.err = d.CallFunction(d.RouterJSFuncName(), d.elementJS)
	// if d.err != "" {
	// 	d.Log(d.err + e)
	// }

	return nil

}

func (d *Dom) setModuleActual(module_name string) (err string) {
	const e = ". setModuleActual"

	// d.Log("CLICK MODULO:", module_name, e)

	if d.actualModule != nil && d.actualModule.ModuleName == module_name {
		return "" // nada que hacer
	}

	d.actualModule, d.err = d.GetModuleByName(module_name)
	if d.err != "" {
		return d.err + e
	}

	return ""
}

var (
	hashOld js.Value
)

func (d *Dom) menuRouter(hashNow js.Value) {
	if hashOld.IsUndefined() {
		hashOld = hashNow
		d.ToggleClass(hashOld, d.MenuClassSelected())
	} else {
		d.ToggleClass(hashOld, d.MenuClassSelected())
		d.ToggleClass(hashNow, d.MenuClassSelected())
		hashOld = hashNow
	}
}
