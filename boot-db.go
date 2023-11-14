package dom

import "github.com/cdvelop/model"

func (d Dom) addBootDataToLocalDB(responses ...model.Response) {

	for _, r := range responses {

		// d.Log("domUpdate .... buscando objeto", r.Object)

		object, err := d.GetObjectByName(r.Object)
		if err != nil {
			d.Log(err)
			continue
		}

		if r.Action == "create" {
			err := d.CreateObjectsInDB(object.Table, false, r.Data)
			if err != nil {
				d.Log("error addBootDataToLocalDB", err)
				continue
			}
		}
	}
}
