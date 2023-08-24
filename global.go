package dom

import (
	"syscall/js"
)

func Call(functionName string, args string) {
	js.Global().Call(functionName, args)
}

func Log(message ...any) {
	js.Global().Get("console").Call("log", message...)
}
