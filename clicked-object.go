package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) objectClickedUI(this js.Value, source_input []js.Value) interface{} {
	const e = ". objectClickedUI error"

	if len(source_input) == 0 || len(source_input) > 2 {
		return d.Log("error se espera: nombre del objeto y opcional id seleccionado" + e)
	}

	d.err = d.setActualObject(source_input[0].String()) //NOMBRE OBJETO
	if d.err != "" {
		return d.Log(d.err + e)
	}

	// d.Log("OBJECTO CLICK:", d.actualObject.ObjectName)

	d.objectID = ""
	if len(source_input) == 2 {
		d.objectID = source_input[1].String() //ID OBJETO
	}

	d.actualObject.FrontHandler.ViewHandlerObject.NotifyStatusChangeAfterClicking()

	if d.actualObject.FrontHandler.AfterClickNotify != nil {

		if d.objectID == "" {

			d.actualObject.FormData = map[string]string{}
			// llamamos al manejador
			d.actualObject.FrontHandler.NotifyUserHasClickedObject()

		} else {

			//1- leer data del objeto
			d.ReadAsyncDataDB(&model.ReadParams{
				FROM_TABLE: d.actualObject.Table,
				ID:         d.objectID,
			}, func(r *model.ReadResults, err string) {

				if err != "" {
					d.Log(err + e)
					return
				}

				// pasamos la data al formulario del objeto
				if len(r.ResultsString) == 1 {
					d.actualObject.FormData = r.ResultsString[0]
				}

				// llamamos al manejador
				d.actualObject.FrontHandler.NotifyUserHasClickedObject()
			})
		}

	} else {
		return d.UserMessage("error objeto:", d.actualObject.ObjectName, "no tiene controlador: NotifyUserHasClickedObject()")
	}

	return nil

}
