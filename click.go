package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) userViewComponentClicked(this js.Value, source_input []js.Value) interface{} {

	if len(source_input) != 2 {
		return d.Log("error se espera: nombre del objeto y id seleccionado")
	}

	object_name := source_input[0].String()
	object_id := source_input[1].String()

	// d.Log("OBJECTO CLICK:", object_name)

	object, err := d.GetObjectByNameMainHandler(object_name)
	if err != "" {
		return d.Log(err)
	}

	if object.FrontHandler.AfterClicked != nil {

		//1- leer data del objeto
		d.ReadAsyncDataDB(model.ReadParams{
			FROM_TABLE: object.Table,
			ID:         object_id,
		}, func(r *model.ReadResults, err string) {

			if err != "" {
				d.Log(err)
				return
			}

			// pasamos la data al formulario del objeto
			if len(r.ResultsString) == 1 {
				object.FormData = r.ResultsString[0]
			}

			// llamamos al manejador
			object.FrontHandler.UserHasClickedObject()

		})

	} else {
		return d.UserMessage("error objeto:", object.ObjectName, "no tiene controlador: UserHasClickedObject(id string)")
	}

	return nil

}
