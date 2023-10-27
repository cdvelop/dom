package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) CallFunction(functionName string, args ...any) error {

	functionValue := js.Global().Get(functionName)
	if !functionValue.Truthy() {
		return model.Error("error la función js:", functionName, "no existe")
	}

	// if !js.Global().Get(functionName).Truthy() {
	// 	return model.Error("la función", functionName, "no existe")
	// }

	result := js.Global().Call(functionName, args...)
	if result.Truthy() { //si retorna algo es por que ocurrió un error
		return model.Error(result.String())
	}

	return nil

}
