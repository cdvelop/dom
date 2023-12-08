package dom

import (
	"strconv"

	"github.com/cdvelop/model"
)

func (d Dom) BuildFrontendUI() (err string) {
	d.Log("CONSTRUYENDO UI...OK")
	const this = "BuildFrontendUI error "

	d.FrontendCheckUser(func(u *model.User, err string) {
		var area string

		if err != "" {
			d.Log(err)
		} else if u != nil {
			area = u.Area
			d.Log("usuario OK:", u.Name)
		}

		menuContainer := doc.Call("querySelector", d.MenuClassName())
		navbarContainer := menuContainer.Get("childNodes").Index(0)
		var index_menu int

		for _, m := range d.GetModules() {

			module_html := doc.Call("querySelector", d.QuerySelectorMenuModule(m.ModuleName))
			if !module_html.IsNull() { // si no es nulo ya existe el modulo el en dom por ende continuamos al siguiente
				continue
			}

			if !m.ModuleSupports(area) { // si no soporta las areas del usuario continuamos al siguiente
				continue
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
				d.Log(this + "usuario no registrado " + err)
			}
		}
	})

	d.registerGlobalFunctions()

	d.Log("UI CONSTRUIDA")
	return
}
