package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (Dom) callFunction(functionName string, args ...any) error {

	if !js.Global().Get(functionName).Truthy() {
		return model.Error("la función", functionName, "no existe")
	}

	js.Global().Call(functionName, args...)

	return nil
}

func Log(message ...any) {
	js.Global().Get("console").Call("log", message...)
}
