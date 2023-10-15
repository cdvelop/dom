package dom

import (
	"github.com/cdvelop/model"
)

func (d Dom) InnerHTML(html_content string, o *model.Object) {

	container, err := d.GetHtmlContainer(o)
	if err != nil {
		d.Log(err)
		return
	}

	container.Set("innerHTML", html_content)

}

func (d Dom) InsertAfterBegin(html_content string, o *model.Object) {
	d.insertInDom("afterbegin", html_content, o)
}

func (d Dom) InsertBeforeEnd(html_content string, o *model.Object) {
	d.insertInDom("beforeend", html_content, o)
}

// where ej: afterbegin,beforeend
func (d Dom) insertInDom(where, html_content string, o *model.Object) {

	container, err := d.GetHtmlContainer(o)
	if err != nil {
		d.Log(err)
		return
	}

	container.Call("insertAdjacentHTML", where, html_content)
}
