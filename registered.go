package dom

import (
	"syscall/js"
)

func (d Dom) registerGlobalFunctions() {

	js.Global().Set("moduleClickedUI", js.FuncOf(d.moduleClickedUI))

	js.Global().Set("objectClickedUI", js.FuncOf(d.objectClickedUI))

	js.Global().Set("deleteObject", js.FuncOf(d.deleteObject))

	js.Global().Set("resetModule", js.FuncOf(d.resetModule))

}
