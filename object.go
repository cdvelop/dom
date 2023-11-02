package dom

import "github.com/cdvelop/model"

func (d *Dom) GetObjectByName(name_to_search string) (*model.Object, error) {
	// d.Log("total objetos:", len(d.objects))
	for _, o := range d.objects {
		// d.Log("BUSCANDO OBJETO:", o.Name)
		if o.Name == name_to_search {
			return o, nil
		}
	}

	return nil, model.Error("error objeto:", name_to_search, ",no encontrado")
}
