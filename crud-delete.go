package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) deleteObject(this js.Value, p []js.Value) interface{} {
	const e = ". deleteObject"
	if len(p) != 2 {
		return d.Log("required 2 args: object name and []ids string " + e)
	}

	object_name := p[0].String() //  arg 1

	d.err = d.SetActualObject(object_name)
	if d.err != "" {
		return d.Log(d.err)
	}

	if d.ObjectActual().FrontHandler.AfterDelete == nil {
		return d.Log("error objeto", d.ObjectActual().ObjectName, "no cuenta con controlador para eliminar")
	}

	object_ids := make([]map[string]string, 0)

	pk_name := d.ObjectActual().PrimaryKeyName()
	// Verificar si el segundo parámetro es un array
	if p[1].Type() == js.TypeObject && p[1].Truthy() && p[1].Get("length").Truthy() {
		length := p[1].Get("length").Int()
		for i := 0; i < length; i++ {

			object_ids = append(object_ids, map[string]string{
				pk_name: p[1].Index(i).String(),
			})

		}
	} else {
		return d.Log("error se esperaba un array de string como parámetro 2" + e)
	}

	d.ReadAsyncDataDB(model.ReadParams{
		FROM_TABLE: d.ObjectActual().Table,
		WHERE:      object_ids,
	}, func(r *model.ReadResults, err string) {

		if err != "" {
			d.Log(err)
			return
		}

		var delete_frontend []map[string]string
		var delete_backend []map[string]string

		for _, item := range r.ResultsString {
			for _, action := range []string{"create", "update", "delete"} {
				if value, exist := item[action]; exist && value != "" {
					// d.Log("-item a eliminar en el frontend:", item)
					delete_frontend = append(delete_frontend, map[string]string{
						pk_name: item[pk_name],
					})
					break
				} else {
					// d.Log("-item a eliminar en el backend:", item)
					delete_backend = append(delete_backend, map[string]string{
						pk_name: item[pk_name],
					})
					break
				}
			}
		}

		// d.Log("delete_frontend:", delete_frontend)
		// d.Log("delete_backend:", delete_backend)

		if len(delete_frontend) != 0 {

			err = d.deleteFrontend(delete_frontend)
			if err != "" {
				d.UserMessage(err)
				return
			}

			d.UserMessage("item eliminado")

		} else if len(delete_backend) != 0 {

			d.SendOneRequest("POST", "delete", object_name, delete_backend, func(resp []map[string]string, err string) {

				if err != "" {
					d.UserMessage(err)
					return
				}

				// d.Log("RESPUESTA SERVIDOR ELIMINACIÓN:", resp)

				err = d.deleteFrontend(delete_backend)
				if err != "" {
					d.UserMessage(err)
					return
				}

				d.UserMessage("item eliminado")

			})

		}

	})

	return nil

}

func (d Dom) deleteFrontend(delete_frontend []map[string]string) (err string) {

	// d.Log("* eliminar en local:", delete_frontend)

	err = d.DeleteObjectsInDB(d.ObjectActual().Table, delete_frontend...)
	if err != "" {
		return err
	}

	if d.ObjectActual().FrontHandler.AfterDelete != nil {
		err = d.ObjectActual().FrontHandler.SetObjectInDomAfterDelete(delete_frontend...)
	}
	return
}
