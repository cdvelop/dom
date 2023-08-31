package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) validateForm(source_input *js.Value) error {

	if d.data_object == nil {
		d.data_object = make(map[string]string, len(d.last_object.Fields))
	}

	// chequear  campos
	for _, field := range d.last_object.Fields {

		input, err := d.getHtmlInputField(field.Name)
		if err != nil {
			return err
		}

		new_value := getHtmlInputValue(field, &input, source_input)

		// log("---new value:", new_value, "campo:", field.Name)

		if field.IsPrimaryKey(d.last_object) && new_value == "" {
			//campo tipo id vació
			continue
		}

		err = inputRight(&field, input, new_value)
		if err != nil {
			return err
		}

		if d.data_object[field.Name] != new_value {

			log("cambio en valor campo:", field.Name, "valor:", new_value)

			d.data_object[field.Name] = new_value
		}

	}

	// var its_update_or_delete bool
	// if d.action_delete || d.action_update {
	// 	its_update_or_delete = true
	// }

	// if d.action_create {
	// 	log("create true")
	// } else {
	// 	log("create false")
	// }

	log("*RESUMEN FORMULARIO:")
	for key, value := range d.data_object {
		log("FIELD NAME: ", key, " VALUE: ", value)
	}

	return nil
}

func (d *Dom) setActionType() {

	id, exist := d.data_object[d.last_object.PrimaryKeyName()]

	if exist {

		log("id existe y no este vació:", id)
		d.action_create = false

		if !d.action_delete {
			d.action_update = true
			log("acción es de tipo update")

		} else {
			log("acción es de tipo delete")
		}

	} else {

		log("no hay id es un objeto nuevo")

		d.action_create = true
		d.action_update = false
		d.action_delete = false
	}

}

func inputRight(field *model.Field, input js.Value, new_value string) error {

	data_option := input.Get("dataset").Get("option").String()

	err := field.Input.Validate.ValidateField(new_value, field.SkipCompletionAllowed, data_option)
	if err == nil {

		// log("value: ", new_value, " input: ", input)

		if new_value != "" {
			js.Global().Call("inputRight", input)
		} else {
			js.Global().Call("inputNormal", input)
		}

		return nil
	}

	js.Global().Call("inputWrong", input, "malo")

	return model.Error("campo", field.Legend, "no valido")
}
