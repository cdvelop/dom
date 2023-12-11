package dom

import (
	"strconv"

	"github.com/cdvelop/model"
)

func (d Dom) BuildFrontendUI() (err string) {
	d.Log("CONSTRUYENDO UI...OK")

	d.FrontendCheckUser(func(u *model.User, err string) {
		var area string

		if u != nil && err == "" {
			area = u.Area

			// AGREGAMOS DATOS DEL USUARIO A LA VISTA
			err = d.InnerHTML(d.QuerySelectorUserName(), u.Name)
			if err != "" {
				d.Log(err)
			}
			err = d.InnerHTML(d.QuerySelectorUserArea(), u.AreaName)
			if err != "" {
				d.Log(err)
			}
		}

		err = d.setUserUI(area)
		if err != "" {
			d.Log(err)
		}

	})

	d.registerGlobalFunctions()

	d.Log("UI CONSTRUIDA")
	return
}

func (d Dom) setUserUI(area string) (err string) {

	menuContainer := doc.Call("querySelector", d.MenuClassName())
	navbarContainer := menuContainer.Get("childNodes").Index(0)

	var index_menu int
	var main_module string
	for _, m := range d.GetModules() {

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
			Content:   m.UI.UserInterface(),
		}

		new_div.Add()

	}

	if area == "" {
		err = d.ElementClicking(d.QuerySelectorMenuModule(d.home_module))
		if err != "" {
			d.Log("en el modulo home " + err)
		}
	} else {

		// click en el modulo principal
		err = d.ElementClicking(d.QuerySelectorMenuModule(main_module))
		if err != "" {
			d.Log("en el modulo principal " + err)
		}
	}

	return
}
