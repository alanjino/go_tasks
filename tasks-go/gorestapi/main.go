package main

import (
	"fmt"
	"log"
	"net/http"
)

type Name struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
	Content string `json:"Content"`
}
type Names []Name

func allNames(w http.ResponseWriter, r *http.Request) {
	a := Names{
		Name{Name: "ALAN", Address: "KANYAKUMARI", Content: "Hello World"},
	}
	fmt.Println("Endpoint hit:All articles endpoint")
	fmt.Fprintf(w, "%s", a)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/names", allNames)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequest()
}
