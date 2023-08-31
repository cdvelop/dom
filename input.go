package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) getHtmlInputField(field_name string) (js.Value, error) {

	field := d.html_form.Get(field_name)
	if field.Truthy() {
		return field, nil
	}
	return js.Value{}, model.Error("campo input html", field_name, "no encontrado")
}

func getHtmlInputValue(fiel model.Field, input, source_input *js.Value) string {

	switch fiel.Input.HtmlName() {
	case "checkbox":
		var out string
		var commaNeeded bool // Variable para indicar si se necesita una coma

		for i := 0; i < input.Length(); i++ {
			check := input.Index(i)
			// log("checkbox", check)
			if check.Get("checked").Bool() {
				if commaNeeded {
					out += ","
				}

				out += check.Get("value").String()
				commaNeeded = true // Ahora necesitas una coma para el siguiente elemento
			}
		}

		*input = *source_input
		return out

	case "radio":
		var out string
		// log("campo de tipo radio")
		for i := 0; i < input.Length(); i++ {
			radio := input.Index(i)
			if radio.Get("checked").Bool() {
				out = radio.Get("value").String()
				break
			}
		}
		*input = *source_input
		return out

	default:
		// log("campo de una sola entrada")
		return input.Get("value").String()

	}
}
