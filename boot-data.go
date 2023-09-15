package dom

func (d Dom) ReadBootDataActions() {

	d.Log("CORRIENDO ACTIONS DATA DE ARRANQUE")

	json := doc.Call("querySelector", "meta[name='JsonBootActions']").Get("content")

	d.Log("json recuperado:", json.String())
}
