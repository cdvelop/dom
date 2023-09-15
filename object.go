package dom

import "github.com/cdvelop/model"

func (d Dom) getObjectByName(name_to_search string) (*model.Object, error) {

	for _, o := range d.objects {

		if o.Name == name_to_search {
			return o, nil
		}
	}

	return nil, model.Error("error objeto:", name_to_search, ",no encontrado")
}
