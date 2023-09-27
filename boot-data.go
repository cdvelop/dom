package dom

func (d Dom) ActionExecutedLater() {

	d.Log("CORRIENDO ACTIONS DATA DE ARRANQUE")

	json := doc.Call("querySelector", "meta[name='JsonBootActions']").Get("content")

	if json.Truthy() {

		resp := d.cut.DecodeResponses([]byte(json.String()))

		d.Log("total respuestas:", len(resp))

		d.domUpdate(resp...)

	}

}
