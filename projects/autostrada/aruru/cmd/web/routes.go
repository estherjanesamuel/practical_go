package main

import (
	"net/http"

	"example.com/assets"

	"github.com/alexedwards/flow"
)

func (app *application) routes() http.Handler {
	mux := flow.New()
	mux.NotFound = http.HandlerFunc(app.notFound)

	mux.Use(app.recoverPanic)
	mux.Use(app.securityHeaders)

	mux.HandleFunc("/", app.home, "GET")

	fileServer := http.FileServer(http.FS(assets.EmbeddedFiles))
	mux.Handle("/static/...", fileServer, "GET")

	return mux
}
