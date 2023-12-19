package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) deleteObject(this js.Value, p []js.Value) interface{} {
	const e = ". deleteObject"
	if len(p) != 2 {
		return d.Log("required 2 args: object name and object id" + e)
	}

	object_name := p[0].String() //  arg 1
	object_id := p[1].String()   // arg 2

	d.err = d.SetActualObject(object_name)
	if d.err != "" {
		return d.Log(d.err)
	}

	if d.ObjectActual().FrontHandler.AfterDelete == nil {
		return d.Log("error objeto", d.ObjectActual().ObjectName, "no cuenta con controlador para eliminar")
	}

	d.Log("ELIMINANDO OBJETO:", d.ObjectActual().ObjectName, "object_id", object_id)

	d.ReadAsyncDataDB(model.ReadParams{
		FROM_TABLE: d.ObjectActual().Table,
		ID:         object_id,
	}, func(r *model.ReadResults, err string) {

		if err != "" {
			d.Log(err)
			return
		}

		for _, data := range r.ResultsString {
			d.Log("DATA:", data)
		}

		if len(r.ResultsString) != 1 {
			d.UserMessage("error", "se esperaba solo un objeto a eliminar")
			return
		}

		d.Log("* id-", object_id, "eliminar en local")
		// err = d.DeleteObjectsInDB(o.Table, data...)
		// if err != nil {
		// 	d.UserMessage("error", err)
		// 	return
		// }
		// Verificar si el objeto existe en el servidor.
		if r.ResultsString[0]["backup"] != "false" {
			d.Log("* id-", object_id, " eliminar en el servidor")

			if d.FetchAdapter == nil {
				d.Log("*error httpAdapter nulo en objeto", d.ObjectActual().ObjectName)
				return
			}

			d.SendOneRequest("POST", "delete", object_name, r.ResultsString, func(resp []map[string]string, err string) {

				if err != "" {
					d.UserMessage(err)
					return
				}

				d.Log("RESPUESTA ELIMINACIÓN:", resp)
				d.UserMessage("elemento eliminado")

			})

		}

	})

	// var data = map[string]interface{}{
	// 	"object_id": object_id,
	// 	"blob":      blob,
	// }

	// d.Log("DATA ANTES DE CREAR:", data)

	// err = d.CreateObjectsInDB(o.Table, true, data)
	// if err != nil {
	// 	return d.Log(err)
	// }
	// // d.Log("DESPUÉS:", data)

	// if o.ObjectViewHandler != nil {

	// 	fiel_id := o.PrimaryKeyName()

	// 	html := o.BuildItemsView(map[string]string{
	// 		fiel_id: data[fiel_id].(string),
	// 		"url":   data["url"].(string),
	// 	})

	// 	d.InsertAfterBegin(html, o)

	// }

	return nil

}
