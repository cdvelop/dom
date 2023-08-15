//go:build js && wasm
// +build js,wasm

package dom

import "syscall/js"

var (
	window, doc, body js.Value
)

func init() {
	window = js.Global()
	doc = window.Get("document")
	body = doc.Get("body")
}
