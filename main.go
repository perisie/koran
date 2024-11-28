package main

import (
	"net/http"
	"perisie.com/koran/src/guice"
	"perisie.com/koran/src/handler"
)

func main() {
	dep := guice.Dep_new(
		"src/template",
		"qurancsv",
		"static",
		"data",
	)
	mux := handler.Mux(dep)
	_ = http.ListenAndServe(":8001", mux)
}
