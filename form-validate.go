package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d *Dom) validateForm(source_input *js.Value) error {

	if d.data_object == nil {
		d.data_object = make(map[string]string, len(d.last_object.Fields))
	}

	// 1 chequear input origen
	source_field_name := source_input.Get("name").String()

	source_field, err := d.last_object.GetFieldByName(source_field_name)
	if err != nil {
		return err
	}

	input, new_value, err := d.getHtmlInput(&source_field)
	if err != nil {
		return err
	}

	err = d.fieldCheck(&source_field, &input, new_value)
	if err != nil {
		return err
	}

	// 2 chequear todos los input menos origen
	for _, field := range d.last_object.Fields {

		if field.Name != source_field_name {

			input, new_value, err := d.getHtmlInput(&field)
			if err != nil {
				return err
			}

			err = d.fieldCheck(&field, &input, new_value)
			if err != nil {
				return err
			}
		}
	}

	log("*RESUMEN FORMULARIO:")
	for key, value := range d.data_object {
		log("FIELD NAME: ", key, " VALUE: ", value)
	}

	return nil
}

func (d *Dom) fieldCheck(field *model.Field, input *js.Value, new_value string) error {

	if field.IsPrimaryKey(d.last_object) && new_value == "" {
		return nil
	}

	err := inputRight(field, *input, new_value)
	if err != nil {
		return err
	}

	if d.data_object[field.Name] != new_value {
		log("---new value:", new_value, "campo:", field.Name)
		d.data_object[field.Name] = new_value
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

	js.Global().Call("inputWrong", input, err.Error())

	return model.Error("campo", field.Legend, "no valido", err.Error())
}
