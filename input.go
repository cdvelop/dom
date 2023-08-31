package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) getHtmlInput(field *model.Field) (input js.Value, value string, err error) {

	input = d.html_form.Get(field.Name)
	if !input.Truthy() {
		return js.Value{}, "", model.Error("input html", field.Name, "no encontrado")
	}
	var temp js.Value

	switch field.Input.HtmlName() {
	case "checkbox":
		var comma bool
		// log("checkbox", field.Name)

		for i := 0; i < input.Length(); i++ {
			check := input.Index(i)
			temp = input.Index(i)

			if check.Get("checked").Bool() {
				if comma {
					value += ","
				}

				value += check.Get("value").String()
				comma = true // se necesita coma para el siguiente elemento
			}
		}

		input = temp

	case "radio":
		// log("campo de tipo radio", field.Name)
		for i := 0; i < input.Length(); i++ {
			radio := input.Index(i)
			temp = input.Index(i)
			if radio.Get("checked").Bool() {
				value = radio.Get("value").String()
				break
			}
		}

		input = temp

	default:
		// log("campo de una sola entrada")
		value = input.Get("value").String()
	}

	return input, value, nil
}
