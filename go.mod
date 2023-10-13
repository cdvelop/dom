module github.com/cdvelop/dom

go 1.20

require (
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/formclient v0.0.1
	github.com/cdvelop/model v0.0.53
	github.com/cdvelop/timeclient v0.0.2
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/timeclient => ../timeclient

replace github.com/cdvelop/formclient => ../formclient

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/indexdb => ../indexdb
