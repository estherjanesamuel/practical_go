package main

import (
	"bank_app/internal/response"
	"net/http"
)

func (app *application) status(w http.ResponseWriter, r *http.Request)  {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w,r,err)
	}
}