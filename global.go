package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) CallFunction(functionName string, args ...any) error {

	if !js.Global().Get(functionName).Truthy() {
		return model.Error("la función", functionName, "no existe")
	}

	js.Global().Call(functionName, args...)

	return nil
}

func (d Dom) Log(message ...any) interface{} {

	for i, msg := range message {
		// Comprueba si el mensaje es de tipo error
		if err, isError := msg.(error); isError {
			// Edita el mensaje y conviértelo a string
			message[i] = err.Error()
		}
	}

	js.Global().Get("console").Call("log", message...)

	return nil
}

func (d Dom) UserMessage(message ...any) interface{} {

	var space string

	var opt = []interface{}{
		"",
	}

	for _, msg := range message {
		// Comprueba si el mensaje es de tipo error
		if err, isError := msg.(error); isError {

			opt[0] = opt[0].(string) + space + err.Error()

			opt = append(opt, "err")

			// Comprueba si el mensaje es de tipo string
		} else if textNew, isString := msg.(string); isString {

			switch textNew {
			case "del", "perm", "stop", "err", "error":
				opt = append(opt, textNew)
			default:

				opt[0] = opt[0].(string) + space + textNew
			}
		}

		space = " "
	}

	err := d.CallFunction(d.h.THEME.FunctionMessageName(), opt...)

	if err != nil {
		d.Log(err)
	}

	return nil
}
