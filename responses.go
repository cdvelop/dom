package dom

import "github.com/cdvelop/model"

func (d Dom) domUpdate(responses ...model.Response) {

	for _, r := range responses {

		d.Log("buscando objeto", r.Object)

		object, err := d.getObjectByName(r.Object)
		if err != nil {
			d.Log(err)
			continue
		}

		switch r.Action {
		case "create":

			if object.FrontendHandler.AfterCreate != nil {

				module_html, err := getHtmlModule(object.ModuleName)
				if err != nil {
					d.Log(err)
					continue
				}

				container_id, tags := object.FrontendHandler.AfterCreate.SetObjectInDomAfterCreate(r.Data...)

				html_container := module_html.Call("querySelector", `div[data-id="`+container_id+`"]`)
				if !html_container.Truthy() {
					d.Log("error no se logro obtener contenedor data-id: " + container_id)
					continue
				}

				// html_container.Set("insertBefore", tags)

				html_container.Call("insertAdjacentHTML", "beforeend", tags)

			}

		case "read":
		case "update":
		case "delete":
		case "error":

		}

	}

}
