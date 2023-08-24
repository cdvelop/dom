package dom

import "strconv"

func (d Dom) BuildUI() {

	Log("¡Hi 7 Go y WebAssembly!")

	d.buildMenu()

	d.buildModules()

	d.registerGlobalFunctions()
}

func (d Dom) buildMenu() {

	menuContainer := doc.Call("querySelector", d.theme.MenuClassName())
	navbarContainer := menuContainer.Get("childNodes").Index(0)

	var index_menu int
	for _, m := range d.modules {
		index_menu++

		li := HtmlElement{
			Container: navbarContainer,
			Name:      "li",
			Class:     d.theme.MenuItemClass(),
			Content:   d.theme.MenuButtonTemplate(m.ModuleName, strconv.Itoa(index_menu), m.IconID, m.Title),
		}

		li.Add()

	}
}

func (d Dom) buildModules() {

	for _, m := range d.modules {

		div := HtmlElement{
			Container: body,
			Name:      "div",
			Id:        m.ModuleName,
			Class:     d.theme.ModuleClassName(),
			Content:   m.UI.UserInterface(),
		}

		div.Add()
	}
}
