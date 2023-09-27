package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func callFunction(functionName string, args ...any) error {

	if !js.Global().Get(functionName).Truthy() {
		return model.Error("la función", functionName, "no existe")
	}

	js.Global().Call(functionName, args...)

	return nil
}

func (d Dom) Log(message ...any) {
	js.Global().Get("console").Call("log", message...)
}

func (d Dom) UserMessage(text string, options ...string) {
	// func (d Dom) message(r model.Response) {

	var opt = []interface{}{
		text,
	}

	for _, o := range options {
		opt = append(opt, o)
	}

	err := callFunction(d.h.THEME.FunctionMessageName(), opt...)

	if err != nil {
		d.Log(err)
	}

}
