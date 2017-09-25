package githubapi

import (
	"encoding/json"
	"fmt"
	"sort"
)

func FormatPrimaryJson(data []byte) Primary {
	var jsontype Primary
	jsonError := json.Unmarshal(data, &jsontype)
	if jsonError != nil {
		fmt.Printf("Failed to parse json: %v\n", jsonError)
	}
	return jsontype
}

func FormatLanguagesJson(data []byte) Languages {

	var jsontype map[string]interface{}
	jsonError := json.Unmarshal(data, &jsontype)
	langList := Languages{}

	if jsonError != nil {
		fmt.Printf("Failed to parse json: %v\n", jsonError)
	}

	// Sort based on: http://ispycode.com/GO/Sorting/Sort-map-by-value
	hack := map[int]string{}
	hackkeys := []int{}

	for key, val := range jsontype {
		hack[int(val.(float64))] = key
		hackkeys = append(hackkeys, int(val.(float64)))
	}
	sort.Ints(hackkeys)

	for i := len(hackkeys) - 1; i >= 0; i-- {
		langList.Languages = append(langList.Languages, hack[hackkeys[i]])
	}

	return langList
}

func FormatCommitterJson(data []byte) Committer {
	fmt.Printf("%v", string(data))
	var jsontype []Committer
	jsonError := json.Unmarshal(data, &jsontype)
	if jsonError != nil {
		fmt.Printf("Failed to parse json: %v\n", jsonError)
	}
	return jsontype[0]
}

func CreateErrorCode(code int, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}
