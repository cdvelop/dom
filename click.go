package dom

import (
	"syscall/js"
)

func (d Dom) ClickModule(module string) {

	for _, m := range d.modules {
		if m.ModuleName == module {
			menuButton := doc.Call("querySelector", d.h.THEME.MenuClassName()+" a[name='"+module+"']")
			if !menuButton.IsUndefined() {
				delayedClick(menuButton)
			}
		}
	}
}

func delayedClick(button js.Value) {
	js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		button.Call("click")
		return nil
	}), 1000)
}
