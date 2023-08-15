//go:build js && wasm
// +build js,wasm

package dom

import "syscall/js"

func Log(message string) {
	js.Global().Get("console").Call("log", message)

}
