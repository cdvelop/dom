package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) UserViewComponentClicked(this js.Value, source_input []js.Value) interface{} {

	if len(source_input) != 2 {
		return d.Log("error se espera: nombre del objeto y id seleccionado")
	}

	object_name := source_input[0].String()
	object_id := source_input[1].String()

	// d.Log("OBJECTO CLICK:", object_name)

	object, err := d.GetObjectByName(object_name)
	if err != "" {
		return d.Log(err)
	}

	if object.FrontHandler.AfterClicked != nil {

		//1- leer data del objeto
		d.ReadStringDataAsyncInDB(model.ReadDBParams{
			FROM_TABLE: object.Table,
			ID:         object_id,
			// WHERE:           []string{object.PrimaryKeyName()},
			// SEARCH_ARGUMENT: object_id,
			// ORDER_BY:        "",
			// SORT_DESC:       false,
		}, func(object_data []map[string]string, err string) {

			if err != "" {
				d.Log(err)
				return
			}

			for _, data := range object_data {
				object.FrontHandler.UserClicked(data)
			}
		})

	} else {
		return d.UserMessage("error", "objeto:", object.ObjectName, "no tiene controlador: UserClicked(id string) error")
	}

	return nil

}
