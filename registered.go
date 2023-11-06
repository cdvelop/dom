package dom

import (
	"syscall/js"
)

func (d *Dom) registerGlobalFunctions() {

	js.Global().Set("userFormTyping", js.FuncOf(d.FormClient.UserFormTyping))

	js.Global().Set("userViewComponentClicked", js.FuncOf(d.UserViewComponentClicked))

	js.Global().Set("saveBlobFile", js.FuncOf(d.saveBlobFile))

	js.Global().Set("deleteObject", js.FuncOf(d.deleteObject))

}
