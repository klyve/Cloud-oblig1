package githubapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Initialize the api endpoint
func Initialize(router *mux.Router) {
	router.HandleFunc("/{username}/{repository}", FindRepository).Methods("GET")
}

// FindRepository http handler
func FindRepository(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	response := FetchAllJSONData(vars["username"], vars["repository"], FetchJSONData)

	output, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		ReturnErrorCode(404, "Not found", w)
		return
	}
	fmt.Fprintf(w, string(output))
}
