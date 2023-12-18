package dom

import "github.com/cdvelop/model"

func (d *Dom) SetActualModule(module_name string) (err string) {

	if d.moduleActual != nil && d.moduleActual.ModuleName == module_name {
		return
	}

	d.moduleActual, err = d.GetModuleByName(module_name)

	return
}

// obtener modulo en uso actualmente por el usuario
func (d *Dom) GetActualModule() (o *model.Module, err string) {
	if d.moduleActual == nil {
		return nil, "modulo actual no definido. GetActualModule"
	}
	return d.moduleActual, ""
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

func (d *Dom) SetActualObject(object_name string) (err string) {
	const e = ". SetActualObject"
	if d.moduleActual == nil {
		return "modulo actual no definido" + e
	}

	err = d.moduleActual.SetActualModuleObject(object_name)
	if err != "" {
		err = err + e
	}

	return
}

func (d *Dom) ObjectActual() *model.Object {

	d.objectActual, d.err = d.moduleActual.GetActualModuleObject()
	if d.err != "" {
		d.Log("ObjectActual", d.err)
		return nil
	}

	return d.objectActual
}

func (d *Dom) GetAllObjects() []*model.Object {

	if len(d.objects) != 0 {
		return d.objects
	}

	for _, m := range d.GetModules() {
		d.objects = append(d.objects, m.Objects...)
	}

	return d.objects
}
