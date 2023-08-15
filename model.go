//go:build js && wasm
// +build js,wasm

package dom

import "github.com/cdvelop/model"

func New(t model.Theme, m ...*model.Module) *Dom {
	return &Dom{
		Theme:   t,
		modules: m,
	}
}

type Dom struct {
	model.Theme
	modules []*model.Module
}
