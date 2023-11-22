package dom

func (d Dom) RunBootActions() {

	// d.Log("CORRIENDO ACTIONS DATA DE ARRANQUE")

	meta := doc.Call("querySelector", "meta[name='JsonBootActions']")
	if !meta.Truthy() {
		return
	}

	json := meta.Get("content")

	if json.Truthy() {

		resp, err := d.DecodeResponses([]byte(json.String()))
		if err != nil {
			d.Log("RunBootActions error", err)
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

func (d Dom) RunTests() {
	// d.Log("EJECUTANDO TEST...")
	meta := doc.Call("querySelector", "meta[name='JsonBootTests']")
	if !meta.Truthy() {
		return
	}

	json := meta.Get("content")

	if json.Truthy() {

		resp, err := d.DecodeResponses([]byte(json.String()))
		if err != nil {
			d.Log("RunTests error", err)
			return
		}

		e := d.Test.RunModuleTests(resp...)
		if e != "" {
			d.Log("RunModuleTests", e)
			return
		}

	}

}
