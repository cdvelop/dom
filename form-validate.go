package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) validateForm(input js.Value) error {

	// Obtener el valor del input pasado desde JavaScript
	input_value := input.Get("value").String()
	input_name := input.Get("name").String()

	// Log(" CAMPO: " + input_name + " ID:" + input_id + " VALOR:" + input_value)

	// 1 validar solo el campo actual
	field, err := d.last_object.GetFieldByName(input_name)
	if err != nil {
		return err
	}

	if !InputRight(field, input, input_value) {
		return model.Error("campo", input_name, "no valido", input_value)
	}

	d.formData()
	// Log("campo : " + input_name + " valido " + input_value)

	//2 validar todo el formulario
	// for _, f := range d.last_object.Fields {

	return nil
}

func InputRight(field model.Field, input js.Value, value string) bool {

	data_option := input.Get("dataset").Get("option").String()

	if field.Input.Validate.ValidateField(value, field.SkipCompletionAllowed, data_option) {

		js.Global().Call("inputRight", input)

		return true
	}

	js.Global().Call("inputWrong", input)

	return false
}
