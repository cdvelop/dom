package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) objectClickedUI(this js.Value, source_input []js.Value) interface{} {
	const e = ". objectClickedUI"
	if len(source_input) != 2 {
		return d.Log("error se espera: nombre del objeto y id seleccionado" + e)
	}

	d.stringVAR = source_input[0].String() //NOMBRE OBJETO
	d.objectID = source_input[1].String()  //ID OBJETO

	// d.Log("OBJECTO CLICK:", d.stringVAR)

	d.err = d.SetActualObject(d.stringVAR)
	if d.err != "" {
		return d.Log(d.err + e)
	}

	if d.ObjectActual().FrontHandler.AfterClicked != nil {

		//1- leer data del objeto
		d.ReadAsyncDataDB(model.ReadParams{
			FROM_TABLE: d.ObjectActual().Table,
			ID:         d.objectID,
		}, func(r *model.ReadResults, err string) {

			if err != "" {
				d.Log(err + e)
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
		return d.UserMessage("error objeto:", d.ObjectActual().ObjectName, "no tiene controlador: UserHasClickedObject()")
	}

	return nil

}
