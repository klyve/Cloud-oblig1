package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/klyve/cloud-oblig1/api/github"
)

// Config struct, contains config information
type Config struct {
	Port int
}

// GetHomePage just returns hello world to the user
func GetHomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	Initialize(false)
}

// Initialize init the application
func Initialize(testing bool) {
	router := mux.NewRouter()

	// Get the environment variables
	cfg := Config{Port: 3000}

	p := strconv.Itoa(cfg.Port)
	portAddr := ":" + p

	subRouter := router.PathPrefix("/projectinfo/v1/github.com/").Subrouter()
	githubapi.Initialize(subRouter)
	router.HandleFunc("/", GetHomePage).Methods("GET")
	if !testing {
		log.Fatal(http.ListenAndServe(portAddr, router))
	}
}
