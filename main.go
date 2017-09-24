package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"

	"github.com/klyve/cloud-oblig1/api/github"
)

type Config struct {
	Port int `env:"PORT" envDefault:"3000"`
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Whats up?")
}

func GetHomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	router := mux.NewRouter()

	// Get the environment variables
	cfg := Config{}
	env.Parse(&cfg)

	p := strconv.Itoa(cfg.Port)
	portAddr := ":" + p

	subRouter := router.PathPrefix("/api").Subrouter()
	githubapi.Initialize(subRouter)
	router.HandleFunc("/", GetHomePage).Methods("GET")
	router.HandleFunc("/people", GetPersonEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(portAddr, router))
}
