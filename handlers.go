package dom

import "github.com/cdvelop/model"

// obtener modulo en uso actualmente por el usuario
func (d *Dom) GetActualModule() (o *model.Module, err string) {
	if d.clickedModule == nil {
		return nil, "modulo actual no definido. GetActualModule"
	}
	return d.clickedModule, ""
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

func (d *Dom) setActualObject(object_name string) (err string) {
	const e = ". setActualObject"

	// nada que hacer
	if d.clickedObject != nil && d.clickedObject.ObjectName == object_name {
		return ""
	}

	// busco el objeto primero en el modulo actual
	if d.clickedModule != nil {
		d.clickedObject, d.err = d.clickedModule.GetActualModuleObject()
		// solo comprobar si no hay error puede que el objeto no sea de este modulo
		if d.err == "" && d.clickedObject.ObjectName == object_name {
			return "" //es el objeto correcto
		}
	}

	// busco el objeto en todos los módulos
	for _, m := range d.GetModules() {

		d.clickedObject, d.err = m.GetObject(object_name)
		if d.clickedObject != nil && d.err == "" {
			// actualizar el objeto actual
			d.clickedObject.SetActualObject()

			return ""
		}
	}

	return "no se logro encontrar objeto:" + object_name + e
}

func (d *Dom) ObjectActual() *model.Object {
	return d.clickedObject
}

func (d *Dom) GetAllObjects() []*model.Object {

	if len(d.objects) != 0 {
		return d.objects
	}

	for _, m := range d.GetModules() {
		d.objects = append(d.objects, m.GetObjects()...)
	}

	return d.objects

}
