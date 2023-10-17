package dom

import (
	"syscall/js"
)

func (d *Dom) registerGlobalFunctions() {

	js.Global().Set("userFormTyping", js.FuncOf(d.FormClient.UserFormTyping))

	js.Global().Set("userViewComponentClicked", js.FuncOf(d.UserViewComponentClicked))

}
