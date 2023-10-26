package dom

import "strconv"

func (d Dom) BuildUI() {

	d.buildMenu()

	d.buildModules()

	d.registerGlobalFunctions()

}

func (d Dom) buildMenu() {

	menuContainer := doc.Call("querySelector", d.h.MenuClassName())
	navbarContainer := menuContainer.Get("childNodes").Index(0)

	var index_menu int
	for _, m := range d.modules {
		index_menu++

		li := HtmlElement{
			Container: navbarContainer,
			Name:      "li",
			Class:     d.h.MenuItemClass(),
			Content:   d.h.MenuButtonTemplate(m.ModuleName, strconv.Itoa(index_menu), m.IconID, m.Title),
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
			Class:     d.h.ModuleClassName(),
			Content:   m.UI.UserInterface(),
		}

		div.Add()
	}
}
