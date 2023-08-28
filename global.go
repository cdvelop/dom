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

func log(message ...any) {
	js.Global().Get("console").Call("log", message...)
}

func (d Dom) userMessage(text string, options ...string) {
	// func (d Dom) message(r model.Response) {

	var opt = []interface{}{
		text,
	}

	for _, o := range options {
		opt = append(opt, o)
	}

	err := callFunction(d.theme.FunctionMessageName(), opt...)

	if err != nil {
		log(err)
	}

}
