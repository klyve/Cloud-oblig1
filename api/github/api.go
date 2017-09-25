package githubapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ReadPrimaryFile(path string) Primary {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	jsontype := FormatPrimaryJson(file)
	return jsontype
}
func ReadLanguagesFile(path string) Languages {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	// fmt.Printf("%s\n", string(file))
	langList := FormatLanguagesJson(file)
	return langList
}
func ReadContributorsFile(path string) Committer {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	contributors := FormatCommitterJson(file)

	return contributors
}

// func ReadJsonFile() Response {
// 	prim := ReadPrimaryFile()
// 	lang := ReadLanguagesFile()
// 	contrib := ReadContributorsFile()
// 	return CombineJsonData(prim, lang, contrib)
// }

func FetchAllJsonData(username string, repo string) interface{} {
	json, err := FetchJsonData("https://api.github.com/repos/" + username + "/" + repo)
	if err != nil {
		fmt.Printf("Could not fetch json %v", err)
		return CreateErrorCode(500, "Internal server error")
	}
	json2, err2 := FetchJsonData("https://api.github.com/repos/" + username + "/" + repo + "/languages")
	if err2 != nil {
		fmt.Printf("Could not fetch json %v", err2)
		return CreateErrorCode(500, "Internal server error")
	}
	json3, err3 := FetchJsonData("https://api.github.com/repos/" + username + "/" + repo + "/contributors")
	if err3 != nil {
		fmt.Printf("Could not fetch json contributors %v", err3)
		return CreateErrorCode(500, "Internal server error")
	}
	fmt.Printf("We here")
	prim := FormatPrimaryJson(json)
	lang := FormatLanguagesJson(json2)
	contrib := FormatCommitterJson(json3)
	return CombineJsonData(prim, lang, contrib)
}

func FetchJsonData(url string) ([]byte, interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "Could not fetch data from github api"
	}
	defer resp.Body.Close()
	body, readError := ioutil.ReadAll(resp.Body)
	// fmt.Printf("%v", body)
	if readError != nil {
		return nil, "Could not parse response body"
	}
	var jsontype Limit

	jsonError := json.Unmarshal(body, &jsontype)
	if jsonError != nil {
		return nil, "Could not parse response body"
	}
	if jsontype.Message != "" {
		return nil, "Rate limit exceeded"
	}

	return body, nil
}

func CombineJsonData(p Primary, l Languages, c Committer) Response {
	var res Response
	res.Name = p.Name
	res.Owner = p.Owner.Login
	res.Committer = c.Login
	res.Contributions = c.Contributions
	res.Languages = l.Languages

	return res
}

func ReturnErrorCode(code int, message string, w http.ResponseWriter) {
	outError, oerr := json.MarshalIndent(CreateErrorCode(code, message), "", "     ")
	if oerr != nil {
		fmt.Fprintf(w, "Internal server error code 500")
		return
	}
	fmt.Fprintf(w, string(outError))
	return
}
