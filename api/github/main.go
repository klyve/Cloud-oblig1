package githubapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SomeFunction() {
	fmt.Print("Hello world")
}

// Some comment
func Initialize(router *mux.Router) {
	router.HandleFunc("/hello", SayHello)
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world from github endpoint")
}
