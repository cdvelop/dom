package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) CallFunction(NameJsFunc string, all_params ...any) (result any, err string) {
	// d.Log("LLAMANDO A FUNCION:", function_name)
	fun_result := js.Global().Get(NameJsFunc)
	// result := global.Call(NameJsFunc, args...)
	// d.Log("RESULTADO BÚSQUEDA FUNCION:", fun_result)
	if !fun_result.Truthy() {
		return nil, "la función js: " + NameJsFunc + " no existe"
	}

	var (
		params []interface{}
		opt    *model.CallJsOptions
	)

	for _, param := range all_params {

		switch p := param.(type) {

		case model.CallJsOptions:
			opt = &p

			params = append(params, opt.Params)

		default:
			params = append(params, p)
		}

	}

	jsValue := js.Global().Call(NameJsFunc, params...)
	result = jsValue

	if opt != nil {

		switch {
		case opt.ResultString:
			result = jsValue.String()
		case opt.ResultInt:
			result = jsValue.Int()
		}

	}

	return
}
