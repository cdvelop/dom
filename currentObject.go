package dom

import (
	"strconv"
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) currentObject(input []js.Value) error {

	if len(input) != 1 {
		return model.Error("en currentObject: se esperaban 1 argumentos y se enviaron:", strconv.Itoa(len(input)))
	}

	d.html_form = input[0].Get("form")
	if d.html_form.IsUndefined() {
		return model.Error("en currentObject: no se logro obtener formulario")
	}

	form_id := d.html_form.Get("id").String()

	if d.last_object == nil {
		// log("primer inicio objeto id: " + form_id)
		return d.getObjectByID(form_id)

	} else {

		if d.last_object.ID() != form_id { //objeto ha cambiado
			log("objeto cambio nuevo: " + form_id + ", anterior: " + d.last_object.ID())

			//reset data formulario
			d.data_object = nil

			return d.getObjectByID(form_id)
		}
	}

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
