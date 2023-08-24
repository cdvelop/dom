package dom

import "syscall/js"

func (d Dom) registerGlobalFunctions() {
	js.Global().Set("userTyping", js.FuncOf(d.userTyping))
}
