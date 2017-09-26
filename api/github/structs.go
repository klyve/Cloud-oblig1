package githubapi

// https://api.github.com/repos/git/git
// https://api.github.com/repos/git/git/contributors
// https://api.github.com/repos/git/git/languages

// Primary struct contains the primary repository data
type Primary struct {
	Name  string `json:"name"`
	Owner struct {
		Login string `json:"login"`
	}
}

// Committer contains the information about the committer
type Committer struct {
	Login         string `json:"login"`
	Contributions int    `json:"contributions"`
}

// Languages contains the languages used in the repo
type Languages struct {
	Languages []string
	Bytes     []float64
}

// Response contains the json success response object
type Response struct {
	Name          string   `json:"project"`
	Owner         string   `json:"owner"`
	Committer     string   `json:"committer"`
	Contributions int      `json:"commits"`
	Languages     []string `json:"language"`
}

// Error contains the error code data
type Error struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}

// Limit for checking if the github api limit is up
type Limit struct {
	Message       string `json:"message"`
	Documentation string `json:"documentation_url"`
}
