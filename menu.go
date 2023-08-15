//go:build js && wasm
// +build js,wasm

package dom

import (
	"fmt"
	"strconv"
)

func (d Dom) BuildMenu() {

	menuContainer := doc.Call("querySelector", ".menu-container")
	navbarContainer := menuContainer.Call("querySelector", ".navbar-container")

	if !navbarContainer.IsNull() {
		fmt.Println("Elemento .menu-container encontrado!")

		var index_menu int
		for _, m := range d.modules {
			index_menu++

			// Crear un nuevo elemento li
			newLi := doc.Call("createElement", "li")
			newLi.Set("className", "navbar-item")

			// Agregar contenido al nuevo elemento li
			newLi.Set("innerHTML", d.MenuButtonTemplate(m.ModuleName, strconv.Itoa(index_menu), m.IconID, m.Title))

			// Agregar el nuevo elemento li a la lista existente
			navbarContainer.Call("appendChild", newLi)

		}
	} else {
		fmt.Println("Elemento .menu-container no encontrado.")
	}

}
