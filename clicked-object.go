package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) objectClickedUI(this js.Value, source_input []js.Value) interface{} {
	const e = ". objectClickedUI"
	if len(source_input) != 2 {
		return d.Log("error se espera: nombre del objeto y id seleccionado")
	}

	d.objectNAME = source_input[0].String()
	d.objectID = source_input[1].String()

	// d.Log("OBJECTO CLICK:", objectNAME)

	d.err = d.SetActualObject(d.objectNAME)
	if d.err != "" {
		return d.Log(d.err + e)
	}

	if d.objectActual.FrontHandler.AfterClicked != nil {

		//1- leer data del objeto
		d.ReadAsyncDataDB(model.ReadParams{
			FROM_TABLE: d.objectActual.Table,
			ID:         d.objectID,
		}, func(r *model.ReadResults, err string) {

			if err != "" {
				d.Log(err + e)
				return
			}

			// pasamos la data al formulario del objeto
			if len(r.ResultsString) == 1 {
				d.objectActual.FormData = r.ResultsString[0]
			}

			// llamamos al manejador
			d.objectActual.FrontHandler.UserHasClickedObject()

		})

	} else {
		return d.UserMessage("error objeto:", d.objectActual.ObjectName, "no tiene controlador: UserHasClickedObject(id string)")
	}

	return nil

}
