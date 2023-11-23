package dom

import (
	"syscall/js"
)

func (d Dom) CallFunction(function_name string, args ...any) (err string) {
	// d.Log("LLAMANDO A FUNCION:", function_name)
	fun_result := js.Global().Get(function_name)
	// result := global.Call(function_name, args...)
	// d.Log("RESULTADO BÚSQUEDA FUNCION:", fun_result)
	if !fun_result.Truthy() {
		return "error la función js:" + function_name + "no existe"
	}

	result := js.Global().Call(function_name, args...)
	if result.Truthy() { //si retorna algo es por que ocurrió un error
		return "CallFunction error: " + result.String()
	}

	return ""
}
