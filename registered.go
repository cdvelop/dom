package dom

import "syscall/js"

func (d Dom) registerGlobalFunctions() {
	js.Global().Set("userFormTyping", js.FuncOf(d.userFormTyping))
}
