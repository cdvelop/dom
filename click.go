package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

// querySelector ej: "a[name='xxx']"
func (d Dom) ElementClicking(querySelector string) error {
	element := doc.Call("querySelector", querySelector)
	// d.Log("ELEMENTO CLICK", element)
	if element.Truthy() {
		element.Call("click")
		return nil
	}

	return model.Error("ElementClicking error no se encontró elemento con la consulta", querySelector)
}

// WaitFor espera el número especificado de milisegundos y luego ejecuta la función de retorno de llamada.
func (d Dom) WaitFor(milliseconds int, callback func()) {
	js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		callback() // Llamar a la función de retorno de llamada después de esperar
		return nil
	}), milliseconds)
}

func (d Dom) UserViewComponentClicked(this js.Value, source_input []js.Value) interface{} {

	if len(source_input) != 2 {
		return d.Log("error se espera: nombre del objeto y id seleccionado")
	}

	object_name := source_input[0].String()
	object_id := source_input[1].String()

	// d.Log("OBJECTO CLICK:", object_name)

	object, err := d.GetObjectByName(object_name)
	if err != nil {
		return d.Log(err)
	}

	if object.FrontHandler.AfterClicked != nil {

		//1- leer data del objeto
		d.ReadStringDataAsyncInDB(model.ReadDBParams{
			FROM_TABLE: object.Table,
			ID:         object_id,
			// WHERE:           []string{object.PrimaryKeyName()},
			// SEARCH_ARGUMENT: object_id,
			// ORDER_BY:        "",
			// SORT_DESC:       false,
		}, func(object_data []map[string]string, err error) {

			if err != nil {
				d.Log(err)
				return
			}

			for _, data := range object_data {
				object.FrontHandler.UserClicked(data)
			}
		})

	} else {
		return d.UserMessage("error", "objeto:", object.ObjectName, "no tiene controlador: UserClicked(id string) error")
	}

	return nil

}
