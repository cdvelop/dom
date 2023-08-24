package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) currentObject(p []js.Value) error {

	if len(p) != 2 {
		return model.MyError{Message: "error currentObject: se esperaban 2 argumentos"}
	}

	form_id := p[1].Get("id").String()

	if d.last_object == nil {
		Log("primer inicio objeto id: " + form_id)

		return d.getObjectByID(form_id)

	} else {

		if d.last_object.ID() != form_id { //objeto ha cambiado

			Log("objeto nuevo: " + form_id + ", anterior: " + d.last_object.ID())

			return d.getObjectByID(form_id)
		}
	}

	Log("*OBJETO ACTUAL: " + d.last_object.ID())

	return nil
}

func (d *Dom) getObjectByID(id string) error {

	object, exist := d.objects[id] //id objeto
	if !exist {
		return model.MyError{Message: "error no se encontró objeto id: " + id}
	}
	d.last_object = object

	Log("*OBJETO ACTUAL: " + d.last_object.ID())

	return nil
}
