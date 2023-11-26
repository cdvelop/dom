package dom

func (d Dom) RunBootActions() {

	// d.Log("CORRIENDO ACTIONS DATA DE ARRANQUE")

	meta := doc.Call("querySelector", "meta[name='JsonBootActions']")
	if !meta.Truthy() {
		return
	}

	json := meta.Get("content")

	content := json.String()

	d.Log("CONTENDIDO JSON BOOT ok 2:", content)
	if json.Truthy() && content != "none" {

		resp, err := d.DecodeResponses([]byte(content))
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
