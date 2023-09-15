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

	form_name := d.html_form.Get("name").String()

	if d.last_object == nil {
		// log("primer inicio objeto id: " + form_name)
		return d.setCurrentObject(form_name)

	} else {

		if d.last_object.Name != form_name { //objeto ha cambiado
			d.Log("objeto cambio nuevo: " + form_name + ", anterior: " + d.last_object.Name)

			//reset data formulario
			d.data_object = nil

			return d.setCurrentObject(form_name)
		}
	}

	return nil
}

func (d *Dom) setCurrentObject(object_name string) error {

	object, err := d.getObjectByName(object_name)
	if err != nil {
		return err
	}

	d.last_object = object

	d.Log("*OBJETO ACTUAL:", d.last_object.Name)

	return nil
}
