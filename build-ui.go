package dom

import (
	"strconv"

	"github.com/cdvelop/model"
)

func (d Dom) BuildFrontendUI() (err string) {
	// d.Log("CONSTRUYENDO UI...OK")

	d.FrontendCheckUser(func(u *model.User, err string) {
		var area string

		if u != nil && err == "" {
			area = u.Area

			d.Log(d.InnerHTML(d.QuerySelectorUserName(), u.Name))

			d.Log(d.InnerHTML(d.QuerySelectorUserArea(), u.AreaName))

		}

		d.Log(d.setUserUI(u, area))

	})

	d.registerGlobalFunctions()

	// d.Log("UI CONSTRUIDA")
	return
}

func (d Dom) setUserUI(u *model.User, area string) (err string) {

	menuContainer := doc.Call("querySelector", d.MenuClassName())
	navbarContainer := menuContainer.Get("childNodes").Index(0)

	var index_menu int
	var main_module string
	for _, m := range d.MainHandlerGetModules() {

		module_html := doc.Call("querySelector", d.QuerySelectorMenuModule(m.ModuleName))
		if !module_html.IsNull() { // si no es nulo ya existe el modulo el en dom por ende continuamos al siguiente
			continue
		}

		if !m.ModuleSupports(area) { // si no soporta las areas del usuario continuamos al siguiente
			continue
		}

		// obtenemos nombre modulo principal
		if !m.ModuleSupports("") && main_module == "" {
			main_module = m.ModuleName
		}

		// CONSTRUIMOS MENU

		index_menu++

		new_li := HtmlElement{
			Container: navbarContainer,
			Name:      "li",
			Class:     d.MenuItemClass(),
			Content:   d.MenuButtonTemplate(m.ModuleName, strconv.Itoa(index_menu), m.IconID, m.Title),
		}
		new_li.Add()

		// CONSTRUIMOS CONTENEDOR MODULO HTML

		new_div := HtmlElement{
			Container: body,
			Name:      "div",
			Id:        m.ModuleName,
			Class:     d.ModuleClassName(),
			Content:   m.UI.UserInterface(u),
		}

		new_div.Add()

	}

	if area == "" {
		main_module = d.home_module
	}

	d.Log(d.ElementClicking(d.QuerySelectorMenuModule(main_module)))

	return
}
