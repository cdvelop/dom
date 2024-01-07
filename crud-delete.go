package dom

import (
	"syscall/js"
)

func (d *Dom) deleteObject(this js.Value, p []js.Value) any {
	const e = ". deleteObject error"
	if len(p) != 2 {
		return d.Log("required 2 args: object name and []ids string " + e)
	}

	object_name := p[0].String() //  arg 1

	d.err = d.SetActualObject(object_name)
	if d.err != "" {
		return d.Log(d.err)
	}

	if d.ObjectActual().FrontHandler.AfterDelete == nil {
		return d.Log("objeto", d.ObjectActual().ObjectName, "no cuenta con controlador para eliminar")
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
		return d.Log("se esperaba un array de string como parámetro" + e)
	}

	d.err = d.DeleteObjectsInDB(d.ObjectActual().Table, true, object_ids...)
	if d.err != "" {
		return d.Log(d.err + e)
	}

	if d.ObjectActual().FrontHandler.AfterDelete != nil {
		d.err = d.ObjectActual().FrontHandler.SetObjectInDomAfterDelete(object_ids...)
		if d.err != "" {
			return d.Log(d.err + e)
		}

		d.UserMessage("item eliminado")

	} else {
		d.Log("objeto:", d.ObjectActual().ObjectName, "no cuenta con FrontHandler.AfterDelete"+e)
	}

	return nil

}
