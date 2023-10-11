package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) ClickModule(module string) {

	for _, m := range d.modules {
		if m.ModuleName == module {
			menuButton := doc.Call("querySelector", d.h.THEME.MenuClassName()+" a[name='"+module+"']")
			if !menuButton.IsUndefined() {
				delayedClick(menuButton)
			}
		}
	}
}

func delayedClick(button js.Value) {
	js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		button.Call("click")
		return nil
	}), 1000)
}

func (d Dom) Clicking(o *model.Object) error {

	module_html, err := d.GetHtmlModule(o.ModuleName)
	if err != nil {
		return err
	}

	err = d.CallFunction(o.ViewComponentName()+"Clicking", *module_html)
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

	object, err := d.GetObjectByName(object_name)
	if err != nil {
		return d.Log(err)
	}

	if object.AfterClicked != nil {

		err := object.UserClicked(object_id)
		if err != nil {
			return d.UserMessage(err)
		}

	} else {
		return d.UserMessage("error", "objeto:", object.Name, "no tiene controlador: UserClicked(id string) error")
		//  return d.Log("error objeto:", object.Name, "no tiene controlador: UserClicked(id string) error")
	}
	// err := f.currentObject(source_input)
	// if err != nil {
	// 	f.dom.Log(err)
	// 	return nil
	// }

	return nil

}
