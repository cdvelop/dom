package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) deleteObject(this js.Value, p []js.Value) any {
	const e = ". deleteObject error"

	if len(p) != 2 {
		return d.Log("required 2 args: object name and []ids string " + e)
	}

	d.actualObject, d.err = d.GetObjectBY(p[0].String(), "") //  arg 1
	if d.err != "" {
		return d.Log(e + d.err)
	}

	objects_ids := make([]map[string]string, 0)

	pk_name := d.actualObject.PrimaryKeyName()
	// Verificar si el segundo parámetro es un array
	if p[1].Type() == js.TypeObject && p[1].Truthy() && p[1].Get("length").Truthy() {
		length := p[1].Get("length").Int()
		for i := 0; i < length; i++ {

			objects_ids = append(objects_ids, map[string]string{
				pk_name: p[1].Index(i).String(),
			})

		}
	} else {
		return d.Log("se esperaba un array de string como parámetro" + e)
	}

	d.err = d.ElementsDelete(d.actualObject.ObjectName, true, objects_ids...)
	if d.err != "" {
		return d.UserMessage(d.err + e)
	}

	return d.UserMessage("item eliminado")
}

func (d *Dom) ElementsDelete(object_name string, on_server_too bool, objects_ids ...map[string]string) (err string) {

	const e = "ElementsDelete "

	d.actualObject, d.err = d.GetObjectBY(object_name, "")
	if d.err != "" {
		return e + d.err
	}

	d.err = d.DeleteObjectsInDB(d.actualObject.Table, on_server_too, objects_ids...)
	if d.err != "" {
		return e + d.err
	}

	if d.actualObject.FrontHandler.AfterDelete != nil {
		d.err = d.actualObject.FrontHandler.SetObjectInDomAfterDelete(objects_ids...)
		if d.err != "" {
			return e + d.err
		}
	}

	if d.actualObject.FrontHandler.ViewHandlerObject != nil {

		d.fnJsCall = model.CallJsOptions{
			NameJsFunc: "delete" + d.actualObject.FrontHandler.ViewHandlerName(),
		}

		for _, data := range objects_ids {

			d.fnJsCall.Params = map[string]any{
				"id": data[d.actualObject.PrimaryKeyName()],
			}

			_, err = d.fnJsCall.CallWithEnableAndQueryParams(d.actualObject)
			if err != "" {
				d.Log(e, "error", err+"id:"+data[d.actualObject.PrimaryKeyName()])
			}
		}

	}

	d.err = d.FormClean(d.actualObject.ObjectName)
	if err != "" {
		d.Log(e, "error", err)
	}

	return ""
}
