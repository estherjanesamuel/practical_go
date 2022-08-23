package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

func home(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w, "hello-world\n")
}

func hello(w http.ResponseWriter, req *http.Request)  {
	name := req.URL.Query().Get("name")
	fmt.Fprintf(w, "hello-%s \n", name)
}