package dom

import (
	"syscall/js"
)

func (d Dom) saveBlobFile(this js.Value, p []js.Value) interface{} {

	if len(p) != 3 {
		return d.Log("error arg expected: object name, object id destiny and blob file")
	}

	object_name := p[0].String() //  arg 1
	object_id := p[1].String()   // arg 2
	blob := p[2]                 // arg 3

	o, err := d.GetObjectByName(object_name)
	if err != "" {
		return d.Log(err)
	}

	var data = map[string]interface{}{
		"object_id": object_id,
		"blob":      blob,
	}

	// d.Log("DATA ANTES DE CREAR:", data)

	err = d.CreateObjectsInDB(o.Table, true, data)
	if err != "" {
		return d.Log(err)
	}
	// d.Log("DESPUÉS:", data)

	if o.FrontHandler.ViewAdapter != nil {

		fiel_id := o.PrimaryKeyName()

		html := o.FrontHandler.BuildItemsView(map[string]string{
			fiel_id: data[fiel_id].(string),
			"url":   data["url"].(string),
		})

		err = d.InsertAfterBegin(d.QuerySelectorObject(o.ModuleName, o.ObjectName), html)
		if err != "" {
			d.Log(err)
		}

	}

	return nil

}
