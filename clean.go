package dom

// func (d *Dom) Init() {
// 	// Otros inicializadores

// 	// Registra el evento beforeunload para limpiar los recursos al cerrar la aplicación
// 	window.Call("addEventListener", "beforeunload", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
// 		d.Cleanup()
// 		return nil
// 	}))
// }

// func (d *Dom) Cleanup() {
//     // Elimina el evento d.keyboardCallback
//     doc.Call("removeEventListener", "keydown", d.keyboardFunc)
//     d.keyboardFunc.Release()

//     // Otras liberaciones de recursos si es necesario

//     d.Log("Recursos liberados correctamente.")
// }
