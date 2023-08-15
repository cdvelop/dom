//go:build js && wasm
// +build js,wasm

package dom

func (d Dom) BuildUI() {

	Log("¡Hi 5 Go y WebAssembly!")

	d.BuildMenu()

	d.BuildModules()

}
