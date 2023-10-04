package dom

import "github.com/cdvelop/model"

func (d Dom) addDataToLocalDB(responses ...model.Response) {

	for _, r := range responses {

		// d.Log("domUpdate .... buscando objeto", r.Object)

		object, err := d.getObjectByName(r.Object)
		if err != nil {
			d.Log(err)
			continue
		}

		switch r.Action {
		case "create":

			if object.FrontendHandler.AfterCreate != nil {

				// module_html, err := getHtmlModule(object.ModuleName)
				// if err != nil {
				// 	d.Log(err)
				// 	continue
				// }

				err = object.FrontendHandler.AfterCreate.SetObjectInDomAfterCreate(r.Data...)
				if err != nil {
					d.UserMessage(err.Error(), "err")
					continue
				}

				// html_container := module_html.Call("querySelector", `div[data-id="`+container_id+`"]`)
				// if !html_container.Truthy() {
				// 	d.Log("error no se logro obtener contenedor data-id:", container_id, "objeto:", object.Name)
				// 	continue
				// }

				// html_container.Set("insertBefore", tags)

				// html_container.Call("insertAdjacentHTML", "beforeend", tags)

			} else {
				d.Log("objeto", object.Name, "no contiene AfterCreate")
			}

		case "read":
			d.Log("HANDLER read NO CREADO EN DOM")
		case "update":
			d.Log("HANDLER update NO CREADO EN DOM")
		case "delete":
			d.Log("HANDLER delete NO CREADO EN DOM")
		case "error":
			d.UserMessage(r.Message, "err")

		}

	}

}
