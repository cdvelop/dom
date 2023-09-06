package dom

func (d Dom) ReadBootDataActions() {

	log("CORRIENDO ACTIONS DATA DE ARRANQUE")

	json := doc.Call("querySelector", "meta[name='JsonBootActions']").Get("content")

	log("json recuperado:", json.String())
}
