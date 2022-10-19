package main

import (
	"bank_app/internal/response"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
)

func(app *application) errorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]
	err := response.JSONWithHeaders(w,status,map[string]string{"Error":message}, headers)
	if err != nil {
		app.logger.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func(app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(),debug.Stack())
	app.logger.Print(trace)

	message := "The server encounterd a problem and could not process your request"
	app.errorMessage(w,r,http.StatusInternalServerError, message,nil)
}