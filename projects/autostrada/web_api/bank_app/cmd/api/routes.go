package main

import (
	"net/http"

	"github.com/alexedwards/flow"
)

func (app *application) routes() http.Handler {
	mux := flow.New()
	mux.HandleFunc("/healtz", app.status, "GET")
	return mux
}