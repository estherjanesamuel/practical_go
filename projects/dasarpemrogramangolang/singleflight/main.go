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
)

func main() {
	http.HandleFunc("/api/report/download/", downloadReportHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func downloadReportHandler(w http.ResponseWriter, req *http.Request) {
	reportId := req.URL.Query().Get("reportid")

	// construct the report path
	reportName := fmt.Sprintf("report-%s.txt", reportId)
	path := filepath.Join(os.TempDir(), reportName)

	// if the report is not exists, generate it first.
	// otherwise, immediatelly download it.
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		log.Println("generate report", reportName, path)

		// simulate long runnign process to generate report
		time.Sleep(5 * time.Second)

		f, err := os.Create(path)
		if err != nil {
			f.Close()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		f.Write([]byte("this is a report"))
		f.Close()
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
	contentDisposition := fmt.Sprintf("attachment; filename=%s", reportName)
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
