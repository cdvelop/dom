package dom

import "github.com/cdvelop/model"

// obtener modulo en uso actualmente por el usuario
func (d *Dom) GetActualModule() (o *model.Module, err string) {
	if d.actualModule == nil {
		return nil, "modulo actual no definido. GetActualModule"
	}
	return d.actualModule, ""
}

func (d *Dom) GetModuleByName(module_name string) (m *model.Module, err string) {
	const e = "func GetModuleByName"
	// d.Log("total objetos:", len(d.objects))
	if module_name == "" {
		return nil, "nombre modulo no puede estar vació" + e
	}

	m, err = d.GetActualModule()
	if m != nil && err == "" && m.ModuleName == module_name {
		return m, ""
	}

	for _, m := range d.GetModules() {
		// d.Log("BUSCANDO OBJETO:", o.ObjectName)
		if m.ModuleName == module_name {
			return m, ""
		}
	}

	return nil, "modulo: " + module_name + ", no encontrado" + e
}
