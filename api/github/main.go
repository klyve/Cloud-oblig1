package githubapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Some comment
func Initialize(router *mux.Router) {
	router.HandleFunc("/{username}/{repository}", FindRepository).Methods("GET")
}

func FindRepository(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	response := FetchAllJsonData(vars["username"], vars["repository"])
	output, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		ReturnErrorCode(500, "Internal server error", w)
		return
	}
	fmt.Fprintf(w, string(output))
}
