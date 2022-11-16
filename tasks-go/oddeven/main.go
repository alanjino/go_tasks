package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	a := NewAPIHandler()

	go a.ReadfromChannel()

	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", a.SayHello)
	r.HandleFunc("/number/{intValue}", a.Print)

	fmt.Printf("Server started at :8080....")
	http.ListenAndServe(":8080", r)

}

var odd = make(chan int)
var even = make(chan int)

type APIHandler struct {
	odd  chan (int)
	even chan (int)
}

type EmployeeDetails struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary   uint64 `json:"salary"`
}

func NewAPIHandler() *APIHandler {
	return &APIHandler{odd: make(chan int), even: make(chan int)}
}

func (a *APIHandler) SayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := r.URL.Query()
	w.Write([]byte(fmt.Sprintf("<h1>%s from %s</h1>", vars["name"], value["location"])))
}

func (a *APIHandler) Print(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%s", vars["intValue"])
	val, _ := strconv.Atoi(vars["intValue"])
	switch val % 2 {
	case 0:
		even <- val
	case 1:
		odd <- val
	default:
	}

}

func (a *APIHandler) ReadfromChannel() {
	for {

		select {
		case <-odd:
			fmt.Println("\n given number is odd")
		case <-even:
			fmt.Println("\n given number is even")
		}
	}

}
