package main

import (
	"fmt"
	"net/http"
)

func startHttpServer() {
	http.HandleFunc("/hello", helloFunc)
	http.ListenAndServe("localhost:8080", nil)
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello: path=%s, qs=%s", r.URL.Path, r.URL.RawQuery)
}

func main() {
	startHttpServer()
}
