package dom

func (d Dom) ActionExecutedLater() {

	// d.Log("CORRIENDO ACTIONS DATA DE ARRANQUE")

	meta := doc.Call("querySelector", "meta[name='JsonBootActions']")
	if !meta.Truthy() {
		return
	}

	json := meta.Get("content")

	if json.Truthy() {

		resp, err := d.DecodeResponses([]byte(json.String()))
		if err != nil {
			d.Log("ActionExecutedLater error", err)
			return
		}

		// d.Log("total respuestas:", len(resp))

		d.addBootDataToLocalDB(resp...)

		for _, o := range d.GetObjects() {

			if o.FrontendHandler.NotifyBootData != nil {
				o.NotifyBootDataIsLoaded()
			}
		}

		// Establece el contenido del elemento meta a una cadena vacía
		meta.Set("content", "")

	}

}
