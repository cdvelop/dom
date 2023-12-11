package dom

func (d Dom) InnerHTML(querySelector, html_content string) (err string) {

	container, err := query(querySelector)
	if err != "" {
		return "InnerHTML error " + err
	}

	container.Set("innerHTML", html_content)

	return
}

func (d Dom) InsertAfterBegin(querySelector, html_content string) (err string) {
	return d.insertInDom(querySelector, "afterbegin", html_content)
}

func (d Dom) InsertBeforeEnd(querySelector, html_content string) (err string) {
	return d.insertInDom(querySelector, "beforeend", html_content)
}

// where ej: afterbegin,beforeend
func (d Dom) insertInDom(querySelector, where, html_content string) (err string) {

	container, err := query(querySelector)
	if err != "" {
		return "insertInDom error " + err
	}

	container.Call("insertAdjacentHTML", where, html_content)

	return
}
