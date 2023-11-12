module github.com/cdvelop/dom

go 1.20

require (
	github.com/cdvelop/cutkey v0.0.60
	github.com/cdvelop/formclient v0.0.14
	github.com/cdvelop/model v0.0.67
)

require (
	github.com/cdvelop/fetchclient v0.0.1
	github.com/cdvelop/logclient v0.0.1
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/logclient => ../logclient

replace github.com/cdvelop/timeclient => ../timeclient

replace github.com/cdvelop/formclient => ../formclient

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/indexdb => ../indexdb

replace github.com/cdvelop/fetchclient => ../fetchclient
