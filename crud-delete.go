package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) deleteObject(this js.Value, p []js.Value) interface{} {

	if len(p) != 2 {
		return d.Log("error deleteObject required 2 args: object name and object id to delete")
	}

	object_name := p[0].String() //  arg 1
	object_id := p[1].String()   // arg 2

	o, err := d.GetObjectByName(object_name)
	if err != nil {
		return d.Log(err)
	}

	if o.AfterDelete == nil {
		return d.Log("error objeto", o.ObjectName, "no cuenta con controlador para eliminar")
	}

	d.Log("ELIMINANDO OBJETO:", o.ObjectName, "object_id", object_id)

	d.ReadStringDataAsyncInDB(model.ReadDBParams{
		FROM_TABLE:      o.Table,
		ID:              object_id,
		WHERE:           []string{},
		SEARCH_ARGUMENT: "",
		ORDER_BY:        "",
		SORT_DESC:       false,
	}, func(data []map[string]string, err error) {

		if err != nil {
			d.Log(err)
			return
		}

		for _, data := range data {
			d.Log("DATA:", data)
		}

		if len(data) != 1 {
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
		if data[0]["backup"] != "false" {
			d.Log("* id-", object_id, " eliminar en el servidor")

			if d.FetchAdapter == nil {
				d.Log("*error httpAdapter nulo en objeto", o.ObjectName)
				return
			}

			d.SendOneRequest("POST", "delete", object_name, data, func(resp []map[string]string, err error) {

				if err != nil {
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

	// if o.ViewHandler != nil {

	// 	fiel_id := o.PrimaryKeyName()

	// 	html := o.BuildItemView(map[string]string{
	// 		fiel_id: data[fiel_id].(string),
	// 		"url":   data["url"].(string),
	// 	})

	// 	d.InsertAfterBegin(html, o)

	// }

	return nil

}
