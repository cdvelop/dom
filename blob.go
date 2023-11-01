package dom

import (
	"syscall/js"
)

func (d Dom) addBlobFileInObject(this js.Value, p []js.Value) interface{} {

	if len(p) != 2 {
		return d.Log("error se espera: archivo blob y id objeto destino")
	}

	blob := p[0]      // Blob argumento 1
	object_id := p[1] // Blob argumento 1

	d.Log("BLOB OK,", blob)
	d.Log("to_object_id,", object_id)

	return nil

}
