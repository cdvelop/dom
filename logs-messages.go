package dom

// support: error,string
func (d Dom) UserMessage(message ...any) interface{} {

	var space string

	var opt = []interface{}{
		"",
	}

	for _, msg := range message {
		// Comprueba si el mensaje es de tipo error
		if err, isError := msg.(error); isError {

			opt[0] = opt[0].(string) + space + err.Error()

			opt = append(opt, "err")

			// Comprueba si el mensaje es de tipo string
		} else if textNew, isString := msg.(string); isString {

			switch textNew {
			case "del", "perm", "stop", "err", "error":
				opt = append(opt, textNew)
			default:

				opt[0] = opt[0].(string) + space + textNew
			}
		}

		space = " "
	}

	err := d.CallFunction(d.h.FunctionMessageName(), opt...)

	if err != nil {
		d.Log(err)
	}

	return nil
}
