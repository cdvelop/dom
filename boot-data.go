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

		err := d.FrontendLoadBootData(content)
		d.UserMessage(err)
		d.Log(err)

		// Establece el contenido del elemento meta a una cadena vacía
		meta.Set("content", "")

	}

}

func (d Dom) FrontendLoadBootData(data string) (err string) {
	const this = "FrontendLoadBootData error "

	if data == "" {
		return this + "sin data para cargar al sistema"
	}

	resp, err := d.DecodeResponses([]byte(data))
	if err != "" {
		return this + err
	}

	// d.Log("total respuestas:", len(resp))

	d.addBootDataToLocalDB(resp...)

	for _, o := range d.GetObjects() {

		if o.FrontHandler.NotifyBootData != nil {
			o.FrontHandler.NotifyBootDataIsLoaded()
		}
	}

	return
}
