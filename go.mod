module github.com/cdvelop/dom

go 1.20

require (
	github.com/cdvelop/cutkey v0.0.54
	github.com/cdvelop/formclient v0.0.14
	github.com/cdvelop/model v0.0.63
	github.com/cdvelop/timeclient v0.0.2
)

require (
	github.com/cdvelop/logclient v0.0.1
	github.com/cdvelop/timetools v0.0.4 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/logclient => ../logclient

replace github.com/cdvelop/timeclient => ../timeclient

replace github.com/cdvelop/formclient => ../formclient

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/indexdb => ../indexdb
