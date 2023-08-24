package dom

import (
	"strings"
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) FormAutoFill(o *model.Object) {

	test_data, err := o.TestData(1, true, false)
	if err != nil {
		Log(err)
	}

	err = formComplete(o, test_data[0])
	if err != nil {
		Log(err)
	}

}

func formComplete(o *model.Object, data map[string]string) error {

	if o == nil {
		return model.MyError{Message: "formComplete error objeto nulo"}
	}

	module_html := body.Call("querySelector", "#"+o.ModuleName())
	if module_html.IsNull() {
		return model.MyError{Message: "formComplete error no se logro obtener modulo html"}
	}

	form := module_html.Call("querySelector", "form", "#"+o.ID())
	if form.IsNull() {
		return model.MyError{Message: "formComplete error no se logro obtener formulario"}
	}

	form.Call("reset")

	for _, f := range o.Fields {

		input, err := getHtmlInput(form, f)
		if err != nil {
			return err
		}

		new_value := data[f.Name]

		switch f.Input.HtmlName() {
		case "checkbox":
			// Log("checkbox: ", f.Name, "tamaño", input.Length(), input)

			for i := 0; i < input.Length(); i++ {

				input_check := input.Index(i)

				value := input_check.Get("value").String()

				if strings.Contains(new_value, value) {
					input_check.Set("checked", true)
				}

				// Log("input check:", input_check, "value:", value)
			}

		case "radio":

			for i := 0; i < input.Length(); i++ {

				input_radio := input.Index(i)

				value := input_radio.Get("value").String()

				if value == new_value {
					input_radio.Set("checked", true)
					break
				}
			}

			// Log("SELECCIÓN radio: ", f.Name, input)

		default:

			// Log("SELECCIÓN: ", f.Input.HtmlName(), f.Name, input)

			input.Set("value", new_value)
		}

	}

	return nil
}

func getHtmlInput(form js.Value, f model.Field) (js.Value, error) {

	var input_type string
	var all string

	var input js.Value

	switch f.Input.HtmlName() {
	case "checkbox", "radio":
		input_type = "input[type='" + f.Input.HtmlName() + "']"
		all = "All"
	}

	input = form.Call("querySelector"+all, input_type+"[name='"+f.Name+"']")
	if input.IsNull() {
		return js.Value{}, model.MyError{Message: "error input: " + f.Name + " tipo:" + f.Input.HtmlName() + " no encontrado en formulario"}
	}

	return input, nil
}
