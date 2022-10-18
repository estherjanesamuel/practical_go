package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	chi "github.com/go-chi/chi/v5"
	"golang.org/x/sync/singleflight"
)

var singleflightGroup singleflight.Group

func main() {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/report/download/{reportId}", downloadReportHandler)
		r.Get("/normal", normal)
		r.Get("/singleflight", singleflightHandler)
	})
	// http.HandleFunc("/api/report/download/", downloadReportHandler)
	fmt.Println("starting web server at :8000")
	http.ListenAndServe(":8000", r)
}

func normal(w http.ResponseWriter, req *http.Request)  {
	start := time.Now()
	status, err := externalCall()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("/normal handler request: processing time: %+v | status %q", time.Since(start), status)
	fmt.Fprintf(w, "response Status: %q", status)
}

func singleflightHandler(w http.ResponseWriter, req *http.Request)  {
	start := time.Now()
	v, err, shared := singleflightGroup.Do("apply_sf_key", func() (interface{}, error) {return externalCall()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status := v.(string)
	if shared {
		log.Printf("/singleflight handler request: processing time: %+v | status %q | shared: %t", time.Since(start),status,shared)
	}
	fmt.Fprintf(w, "response Status: %q", status)
}

func externalCall() (string, error) {
	time.Sleep(300 * time.Millisecond)
	log.Println("external call")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response external call: %s", resp.Status)
	}
	log.Println("response external call:", resp.Status)
	return resp.Status, err
}

func downloadReportHandler(w http.ResponseWriter, req *http.Request) {
	reportId := chi.URLParam(req, "reportId")

	// construct the report path
	reportName := fmt.Sprintf("report-%s.txt", reportId)
	path := filepath.Join(os.TempDir(), reportName)


	// implemented singleflight group
	sharedProcessKey := fmt.Sprintf("generate %s", reportName)
	_, err, shared := singleflightGroup.Do(sharedProcessKey, func() (interface{}, error) {
		// if the report is not exist, generate it first.
		// otherwise, immediatelly download it first
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			log.Println("generate report", reportName, path)
	
			// simulate long runnign process to generate report
			time.Sleep(5 * time.Second)
	
			f, err := os.Create(path)
			if err != nil {
				f.Close()
				return nil, err
			}
	
			f.Write([]byte("this is a report"))
			f.Close()
		}

		return true, nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if shared {
		log.Printf("generation of report %v is shared with others", reportName)
	}

	// open the file, download it
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if f != nil {
		defer f.Close()
	}
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set Content-Disposition for tell the browser that this is a report generation operation 
	contentDisposition := fmt.Sprintf("attachment; filename=%s", reportName)
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
