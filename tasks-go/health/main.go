package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
func health(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "<h1> Health Check </h1>\n")
}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}
