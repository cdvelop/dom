package dom

func (d *Dom) setActualObject(object_name string) (err string) {

	d.actualObject, err = d.GetObjectBY(object_name, "")
	if err != "" {
		return "setActualObject: " + err

	}

	d.actualObject.Module.SetActualObject(d.actualObject)
	return
}
