package dom

func (d Dom) userMessage(text string, options ...string) {
	// func (d Dom) message(r model.Response) {

	var opt = []interface{}{
		text,
	}

	for _, o := range options {
		opt = append(opt, o)
	}

	err := d.callFunction(d.theme.FunctionMessageName(), opt...)

	if err != nil {
		Log(err)
	}

}
