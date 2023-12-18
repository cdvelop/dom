package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) objectClickedUI(this js.Value, source_input []js.Value) interface{} {

	if len(source_input) != 2 {
		return d.Log("error se espera: nombre del objeto y id seleccionado")
	}

	d.object_name = source_input[0].String()
	d.object_id = source_input[1].String()

	// d.Log("OBJECTO CLICK:", object_name)

	d.err = d.SetActualObject(d.object_name)
	if d.err != "" {
		return d.Log(d.err)
	}

	if d.ObjectActual().FrontHandler.AfterClicked != nil {

		//1- leer data del objeto
		d.ReadAsyncDataDB(model.ReadParams{
			FROM_TABLE: d.ObjectActual().Table,
			ID:         d.object_id,
		}, func(r *model.ReadResults, err string) {

			if err != "" {
				d.Log(err)
				return
			}

			// pasamos la data al formulario del objeto
			if len(r.ResultsString) == 1 {
				d.ObjectActual().FormData = r.ResultsString[0]
			}

			// llamamos al manejador
			d.ObjectActual().FrontHandler.UserHasClickedObject()

		})

	} else {
		return d.UserMessage("error objeto:", d.ObjectActual().ObjectName, "no tiene controlador: UserHasClickedObject(id string)")
	}

	return nil

}
