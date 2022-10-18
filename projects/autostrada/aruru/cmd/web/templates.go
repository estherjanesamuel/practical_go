package main

import (
	"net/http"

	"example.com/internal/version"
)

func (app *application) newTemplateData(r *http.Request) map[string]any {
	data := map[string]any{
		"Version": version.Get(),
	}

	return data
}
