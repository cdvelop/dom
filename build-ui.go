package dom

import (
	"strconv"
)

func (d Dom) BuildUI() (err string) {
	const this = "BuildUI error "
	var area string
	user, err := d.GetLoginUser(nil)
	if err == "" {
		area = user.Area
		d.Log("usuario:", user.Name)
	}

	d.buildMenu(area)

	d.buildModules(area)

	if user == nil { // usuario no registrado
		err = d.ElementClicking(d.QuerySelectorMenuModule(d.home_module))
		if err != "" {
			return this + err
		}
	}

	d.registerGlobalFunctions()

	d.Log("UI CONSTRUIDA")
	return
}

func (d Dom) buildMenu(area string) {

	menuContainer := doc.Call("querySelector", d.MenuClassName())
	navbarContainer := menuContainer.Get("childNodes").Index(0)

	var index_menu int
	for _, m := range d.GetModules() {
		// d.Log(len(m.Areas), "AREAS DEFINIDAS", m.ModuleName, m.Areas)

		if !m.ModuleSupports(area) {
			continue
		}

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

func (d Dom) buildModules(area string) {

	for _, m := range d.GetModules() {

		if !m.ModuleSupports(area) {
			continue
		}

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
