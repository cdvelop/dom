package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) ClickModule(module string) error {

	for _, m := range d.modules {
		if m.ModuleName == module {
			menuButton := doc.Call("querySelector", d.h.MenuClassName()+" a[name='"+module+"']")
			if !menuButton.IsUndefined() {
				delayed()
				menuButton.Call("click")
			} else {
				return model.Error("no se encontró modulo", module, " en el menu para la acción click")
			}
		}
	}
	return nil
}

func delayed() {
	js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		return nil
	}), 50)
}

func (d Dom) Clicking(o *model.Object, id string) error {

	module_html, err := d.GetHtmlModule(o.ModuleName)
	if err != nil {
		return err
	}

	if o.ViewHandler == nil {
		return model.Error("error objeto", o.Name, "no tiene controlador ViewHandler para realizar click")
	}

	err = d.CallFunction(o.ViewComponentName()+"Clicking", *module_html, id)
	if err != nil {
		return err
	}

	return nil
}

func (d Dom) ClickingNEW(o *model.Object, id string) (err error) {
	var moduleHTML *js.Value
	moduleHTML, err = d.GetHtmlModule(o.ModuleName)
	if err != nil {
		return err
	}

	maxRetries := 10     // Número máximo de reintentos
	retryInterval := 5   // Intervalo de reintento en milisegundos
	maxRetryTime := 2000 // Tiempo máximo de reintento en milisegundos (1 segundo)

	startTime := js.Global().Get("Date").New().Call("getTime").Int()

	for retries := 0; retries < maxRetries; retries++ {
		err = d.CallFunction(o.ViewComponentName()+"Clicking", *moduleHTML, id)
		if err == nil {
			return nil // Éxito, no hay error
		}

		// Si el tiempo de reintento ha excedido el tiempo máximo, salimos del bucle
		currentTime := js.Global().Get("Date").New().Call("getTime").Int()
		if currentTime-startTime > maxRetryTime {
			break
		}

		// Espera antes de intentar nuevamente
		for start := js.Global().Get("Date").New().Call("getTime").Int(); js.Global().Get("Date").New().Call("getTime").Int()-start < retryInterval; {
		}
	}

	return err // Devuelve el error si no se pudo completar después de los reintentos
}

func (d Dom) UserViewComponentClicked(this js.Value, source_input []js.Value) interface{} {

	if len(source_input) != 2 {
		return d.Log("error se espera: nombre del objeto y id seleccionado")
	}

	object_name := source_input[0].String()
	object_id := source_input[1].String()

	object, err := d.GetObjectByName(object_name)
	if err != nil {
		return d.Log(err)
	}

	if object.AfterClicked != nil {

		//1- leer data del objeto
		d.h.ReadDataAsyncInDB(
			object.Table,
			[]map[string]string{{
				"WHERE": object.PrimaryKeyName(),
				"ARGS":  object_id,
			},
			}, func(object_data []map[string]string, err error) {

				if err != nil {
					d.Log(err)
					return
				}

				for _, data := range object_data {
					object.UserClicked(data)
				}
			})

	} else {
		return d.UserMessage("error", "objeto:", object.Name, "no tiene controlador: UserClicked(id string) error")
		//  return d.Log("error objeto:", object.Name, "no tiene controlador: UserClicked(id string) error")
	}
	// err := f.currentObject(source_input)
	// if err != nil {
	// 	f.dom.Log(err)
	// 	return nil
	// }

	return nil

}
