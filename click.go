package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) ClickModule(module_name string) error {

	menuButton := doc.Call("querySelector", d.MenuClassName()+" a[name='"+module_name+"']")
	if !menuButton.IsUndefined() {
		delayed()
		menuButton.Call("click")
	} else {
		return model.Error("modulo", module_name, " no encontrado en el menu para la acción click")
	}

	return nil
}

func delayed() {
	js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		return nil
	}), 50)
}

func (d Dom) Clicking(o *model.Object, id string) error {

	module_html, err := GetHtmlModule(o.ModuleName)
	if err != nil {
		return err
	}

	if o.ViewHandler == nil {
		return model.Error("error objeto", o.ObjectName, "no tiene controlador ViewHandler para realizar click")
	}

	err = d.CallFunction(o.ViewHandlerName()+"Clicking", *module_html, id)
	if err != nil {
		return err
	}

	return nil
}

func (d Dom) UserViewComponentClicked(this js.Value, source_input []js.Value) interface{} {

	if len(source_input) != 2 {
		return d.Log("error se espera: nombre del objeto y id seleccionado")
	}

	object_name := source_input[0].String()
	object_id := source_input[1].String()

	// d.Log("OBJECTO CLICK:", object_name)

	object, err := d.GetObjectByName(object_name)
	if err != nil {
		return d.Log(err)
	}

	if object.AfterClicked != nil {

		//1- leer data del objeto
		d.ReadStringDataAsyncInDB(model.ReadDBParams{
			FROM_TABLE: object.Table,
			ID:         object_id,
			// WHERE:           []string{object.PrimaryKeyName()},
			// SEARCH_ARGUMENT: object_id,
			// ORDER_BY:        "",
			// SORT_DESC:       false,
		}, func(object_data []map[string]string, err error) {

			if err != nil {
				d.Log(err)
				return
			}

			for _, data := range object_data {
				object.UserClicked(data)
			}
		})

	} else {
		return d.UserMessage("error", "objeto:", object.ObjectName, "no tiene controlador: UserClicked(id string) error")
	}

	return nil

}
