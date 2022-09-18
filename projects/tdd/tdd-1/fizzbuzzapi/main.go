package main

import (
	"fizzbuzzcore"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main()  {
	http.HandleFunc("/fizzbuzz", fizzbuzz)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func fizzbuzz(w http.ResponseWriter, req *http.Request)  {
	fizzbuzz := req.URL.Query().Get("n")

	if fizzbuzz == "" {
		fmt.Fprintf(w, "fizzbuzz is empty")
		return
	}
	if n, err := strconv.ParseInt(fizzbuzz,10,32); err != nil {
		fmt.Fprintf(w, "Invalid fizzbuzz!")
		return
	} else {
		for i := 1; i <= int(n); i++ {
			fmt.Fprintf(w, "%d ==> %q \n", i, fizzbuzzcore.GetValue(i))
		}
	}
}