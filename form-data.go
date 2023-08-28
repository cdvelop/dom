package dom

import (
	"syscall/js"
)

func (d *Dom) formData() {

	if d.data_object == nil {
		// creamos la variable par almacenar el formulario si este no existe
		d.data_object = make(map[string]string, len(d.last_object.Fields))
	}

	form_data := js.Global().Get("FormData").New(d.form)

	form_data.Call("forEach", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		new_value := p[0].String()
		field_name := p[1].String()

		if _, exist := d.last_object.FieldExist(field_name); exist {

			if existing_values, ok := d.data_object[field_name]; ok {

				d.data_object[field_name] = existing_values + ", " + new_value

			} else {
				d.data_object[field_name] = new_value
			}

		}

		return nil
	}), nil)

	for key, value := range d.data_object {
		log("FIELD NAME: ", key, "VALUE: ", value)
	}

}
