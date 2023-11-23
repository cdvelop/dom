package dom

func (d Dom) RunTests() {
	// d.Log("EJECUTANDO TEST...")
	meta := doc.Call("querySelector", "meta[name='JsonBootTests']")
	if !meta.Truthy() {
		return
	}

	json := meta.Get("content")

	if json.Truthy() {

		resp, err := d.DecodeResponses([]byte(json.String()))
		if err != "" {
			d.Log("RunTests error", err)
			return
		}

		d.Test.RunModuleTests(resp, func(err string) {

			if err != "" {
				d.Log("error", "RunModuleTests", err)
			} else {
				d.Log("Resultados Pruebas ok")
			}
		})

	}

}
