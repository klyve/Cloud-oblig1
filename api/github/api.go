package githubapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// ReadPrimaryFile reads the primary file for testing
func ReadPrimaryFile(path string) Primary {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	jsontype := FormatPrimaryJSON(file)
	return jsontype
}

// ReadLanguagesFile reads the languages file for testing
func ReadLanguagesFile(path string) Languages {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	// fmt.Printf("%s\n", string(file))
	langList := FormatLanguagesJSON(file)
	return langList
}

// ReadContributorsFile reads the contributors file for testing
func ReadContributorsFile(path string) Committer {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	contributors := FormatCommitterJSON(file)

	return contributors
}

// FetchJSONFunc type for Dependency injection
type FetchJSONFunc func(string) ([]byte, interface{})

// FetchAllJSONData fetches the json data and combines it
func FetchAllJSONData(username string, repo string, fetcData FetchJSONFunc) interface{} {

	json, err := fetcData("https://api.github.com/repos/" + username + "/" + repo)
	if err != nil {
		fmt.Printf("Could not fetch json %v", err)
		return CreateErrorCode(500, "Internal server error")
	}
	json2, err2 := fetcData("https://api.github.com/repos/" + username + "/" + repo + "/languages")
	if err2 != nil {
		fmt.Printf("Could not fetch json %v", err2)
		return CreateErrorCode(500, "Internal server error")
	}
	json3, err3 := fetcData("https://api.github.com/repos/" + username + "/" + repo + "/contributors")
	if err3 != nil {
		fmt.Printf("Could not fetch json contributors %v", err3)
		return CreateErrorCode(500, "Internal server error")
	}
	prim := FormatPrimaryJSON(json)
	lang := FormatLanguagesJSON(json2)
	contrib := FormatCommitterJSON(json3)
	// contrib := ReadContributorsFile("./api/github/json/contributors.json")
	return CombineJSONData(prim, lang, contrib)
}

// FetchJSONData fetches the json data from the web
func FetchJSONData(url string) ([]byte, interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "Could not fetch data from github api"
	}
	// defer resp.Body.Close()
	body, readError := ioutil.ReadAll(resp.Body)
	// fmt.Printf("%v", body)
	if readError != nil {
		return nil, "Could not parse response body"
	}
	var jsontype Limit

	jsonError := json.Unmarshal(body, &jsontype)
	if jsonError != nil {
		return body, nil
	}
	if jsontype.Message != "" {
		return nil, "Rate limit exceeded"
	}

	return body, nil
}

// CombineJSONData combines the data into a struct
func CombineJSONData(p Primary, l Languages, c Committer) Response {
	var res Response
	res.Name = p.Name
	res.Owner = p.Owner.Login
	res.Committer = c.Login
	res.Contributions = c.Contributions
	res.Languages = l.Languages

	return res
}

// ReturnErrorCode returns an error code
func ReturnErrorCode(code int, message string, w http.ResponseWriter) {
	outError, oerr := json.MarshalIndent(CreateErrorCode(code, message), "", "     ")
	if oerr != nil {
		fmt.Fprintf(w, "Internal server error code 500")
		return
	}
	fmt.Fprintf(w, string(outError))
	return
}
