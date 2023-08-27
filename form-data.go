package dom

import (
	"syscall/js"
)

func (d Dom) formData() map[string]string {

	form_data := js.Global().Get("FormData").New(d.form)
	data_object := make(map[string]string)

	form_data.Call("forEach", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		new_value := p[0].String()
		field_name := p[1].String()

		if _, exist := d.last_object.FieldExist(field_name); exist {

			if existing_values, ok := data_object[field_name]; ok {

				data_object[field_name] = existing_values + ", " + new_value

			} else {
				data_object[field_name] = new_value
			}

		}

		return nil
	}), nil)

	for key, value := range data_object {
		Log("FIELD NAME: ", key, "VALUE: ", value)

	}

	return data_object
}
