package dom

import "strconv"

func (d Dom) BuildUI() {

	d.buildMenu()

	d.buildModules()

	d.registerGlobalFunctions()

	d.Log("UI CONSTRUIDA")
}

func (d Dom) buildMenu() {

	menuContainer := doc.Call("querySelector", d.MenuClassName())
	navbarContainer := menuContainer.Get("childNodes").Index(0)

	var index_menu int
	for _, m := range d.GetModules() {
		index_menu++

		li := HtmlElement{
			Container: navbarContainer,
			Name:      "li",
			Class:     d.MenuItemClass(),
			Content:   d.MenuButtonTemplate(m.ModuleName, strconv.Itoa(index_menu), m.IconID, m.Title),
		}

		li.Add()

	}
}

func (d Dom) buildModules() {

	for _, m := range d.GetModules() {

		div := HtmlElement{
			Container: body,
			Name:      "div",
			Id:        m.ModuleName,
			Class:     d.ModuleClassName(),
			Content:   m.UI.UserInterface(),
		}

		div.Add()
	}
}
