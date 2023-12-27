package dom

import "syscall/js"

func (d Dom) ToggleClass(elem js.Value, className string) {
	d.elementJS = elem.Get("classList")
	d.elementJS.Call("toggle", className)
}
