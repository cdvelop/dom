//go:build js && wasm
// +build js,wasm

package dom

import (
	"syscall/js"
)

func Call(functionName string, args string) {
	js.Global().Call(functionName, args)
}
