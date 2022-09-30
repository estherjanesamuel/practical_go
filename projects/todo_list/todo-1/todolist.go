package main

import (
	"io"
	"net/http"
	log "github.com/sirupsen/logrus"
	chi "github.com/go-chi/chi/v5"
)


func main()  {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/healthz", Healthz)
	})
	
	port := ":8000"
	log.Info("Starting Todolist API Server on Port ", port)
	http.ListenAndServe(port, r)
}

func Healthz(w http.ResponseWriter, r *http.Request)  {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func init()  {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}