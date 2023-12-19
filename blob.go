package dom

import (
	"syscall/js"
)

func (d *Dom) saveBlobFile(this js.Value, p []js.Value) interface{} {
	const e = ". saveBlobFile"
	if len(p) != 3 {
		return d.Log("error arg expected: object name, object id destiny and blob file")
	}

	object_name := p[0].String() //  arg 1
	object_id := p[1].String()   // arg 2
	blob := p[2]                 // arg 3

	d.err = d.SetActualObject(object_name)
	if d.err != "" {
		return d.Log(d.err + e)
	}

	var data = map[string]interface{}{
		"object_id": object_id,
		"blob":      blob,
	}

	// d.Log("DATA ANTES DE CREAR:", data)

	d.err = d.CreateObjectsInDB(d.ObjectActual().Table, true, data)
	if d.err != "" {
		return d.Log(d.err + e)
	}
	// d.Log("DESPUÉS:", data)

	if d.ObjectActual().FrontHandler.ObjectViewHandler != nil {

		fiel_id := d.ObjectActual().PrimaryKeyName()

		html := d.ObjectActual().FrontHandler.BuildItemsView(map[string]string{
			fiel_id: data[fiel_id].(string),
			"url":   data["url"].(string),
		})

		d.err = d.InsertAfterBegin(d.QuerySelectorObject(d.ObjectActual().ModuleName, d.ObjectActual().ObjectName), html)
		if d.err != "" {
			d.Log(d.err + e)
		}

	}

	return nil

}
