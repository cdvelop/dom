package dom

import (
	"strconv"
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) currentObject(p []js.Value) error {

	if len(p) != 1 {
		return model.Error("en currentObject: se esperaban 1 argumentos y se enviaron:", strconv.Itoa(len(p)))
	}

	d.form = p[0].Get("form")
	if d.form.IsUndefined() {
		return model.Error("en currentObject: no se logro obtener formulario")
	}

	form_id := d.form.Get("id").String()

	if d.last_object == nil {
		// log("primer inicio objeto id: " + form_id)
		return d.getObjectByID(form_id)

	} else {

		if d.last_object.ID() != form_id { //objeto ha cambiado
			log("objeto cambio nuevo: " + form_id + ", anterior: " + d.last_object.ID())
			return d.getObjectByID(form_id)
		}
	}

	log("*OBJETO ACTUAL: "+d.last_object.ID(), "form ok")

	return nil
}

func (d *Dom) getObjectByID(id string) error {

	object, exist := d.objects[id] //id objeto
	if !exist {
		return model.Error("error no se encontró objeto id:", id)
	}
	d.last_object = object

	log("*OBJETO ACTUAL: " + d.last_object.ID())

	return nil
}
