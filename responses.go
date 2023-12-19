package dom

import "github.com/cdvelop/model"

func (d *Dom) addDataToLocalDB(responses ...model.Response) {
	const e = ". addDataToLocalDB"
	for _, r := range responses {

		// d.Log("domUpdate .... buscando objeto", r.Object)

		d.err = d.setActualObject(r.Object)
		if d.err != "" {
			d.Log(d.err + e)
			continue
		}

		switch r.Action {
		case "create":

			if d.ObjectActual().FrontHandler.AfterCreate != nil {

				// module_html, err := getHtmlModule(d.ObjectActual().ModuleName)
				// if err != nil {
				// 	d.Log(err)
				// 	continue
				// }

				d.err = d.ObjectActual().FrontHandler.AfterCreate.SetObjectInDomAfterCreate(r.Data...)
				if d.err != "" {
					d.UserMessage("error", d.err)
					continue
				}

				// html_container := module_html.Call("querySelector", `div[data-id="`+container_id+`"]`)
				// if !html_container.Truthy() {
				// 	d.Log("error no se logro obtener contenedor data-id:", container_id, "objeto:", d.ObjectActual().Name)
				// 	continue
				// }

				// html_container.Set("insertBefore", tags)

				// html_container.Call("insertAdjacentHTML", "beforeend", tags)

			} else {
				d.Log("objeto", d.ObjectActual().ObjectName, "no contiene AfterCreate")
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
