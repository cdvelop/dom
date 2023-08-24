package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) validateForm(input, form js.Value) error {

	// Obtener el valor del input pasado desde JavaScript
	input_value := input.Get("value").String()
	input_name := input.Get("name").String()
	// input_id := input.Get("id").String()
	// fieldValue := inputValue.String()

	// Log(" CAMPO: " + input_name + " ID:" + input_id + " VALOR:" + input_value)

	// 1 validar solo el campo actual
	field := d.last_object.GetFieldByName(input_name)
	if !InputRight(field, input, input_value) {
		return model.MyError{Message: "campo " + input_name + " no valido " + input_value}
	}

	d.formData(form)
	// Log("campo : " + input_name + " valido " + input_value)

	//2 validar todo el formulario
	// for _, f := range d.last_object.Fields {

	return nil
}

func InputRight(field model.Field, input js.Value, value string) bool {

	if field.Input.Validate.ValidateField(value, field.SkipCompletionAllowed) {

		js.Global().Call("InputRight", input)

		return true
	}

	js.Global().Call("InputWrong", input)

	return false
}
