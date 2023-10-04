package dom

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (d Dom) InsertAfterBegin(vHandlers ...model.ViewHandler) {
	d.insertInDom("afterbegin", vHandlers...)
}

func (d Dom) InsertBeforeEnd(vHandlers ...model.ViewHandler) {
	d.insertInDom("beforeend", vHandlers...)
}

// where ej: afterbegin,beforeend
func (d Dom) insertInDom(where string, vHandlers ...model.ViewHandler) {
	var container js.Value

	for _, v := range vHandlers {

		if !container.Truthy() {
			container = doc.Call("querySelector", "div#"+v.ObjectVIEW().ModuleName+" div[data-id='"+v.ObjectVIEW().Name+"']")
			if !container.Truthy() {
				d.Log("error no se logro obtener contenedor objeto:", v.ObjectVIEW().Name)
				break
			}
		}

		container.Call("insertAdjacentHTML", where, v.BuildTag())

		// d.Log("InsertAfterBegin Object:", v.ObjectVIEW().Name)
	}
	// d.Log("contenedor:", container)
}
