package dom

func (d Dom) RunBootActions() {

	// d.Log("CORRIENDO ACTIONS DATA DE ARRANQUE")

	meta := doc.Call("querySelector", "meta[name='JsonBootActions']")
	if !meta.Truthy() {
		return
	}

	json := meta.Get("content")

	d.Log("CONTENDIDO JSON BOOT ok 1:", json.String())
	if json.Truthy() {

		resp, err := d.DecodeResponses([]byte(json.String()))
		if err != "" {
			d.Log("RunBootActions ", err)
			return
		}

		// d.Log("total respuestas:", len(resp))

		d.addBootDataToLocalDB(resp...)

		for _, o := range d.GetObjects() {

			if o.FrontHandler.NotifyBootData != nil {
				o.FrontHandler.NotifyBootDataIsLoaded()
			}
		}

		// Establece el contenido del elemento meta a una cadena vacía
		meta.Set("content", "")

	}

}
