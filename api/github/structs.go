package githubapi

// https://api.github.com/repos/git/git
// https://api.github.com/repos/git/git/contributors
// https://api.github.com/repos/git/git/languages

type Primary struct {
	Name  string `json:"name"`
	Owner struct {
		Login string `json:"login"`
	}
}

type Committer struct {
	Login         string `json:"login"`
	Contributions int    `json:"contributions"`
}

type Languages struct {
	Languages []string
	Bytes     []float64
}

type Response struct {
	Name          string   `json:"name"`
	Owner         string   `json:"owner"`
	Committer     string   `json:"committer"`
	Contributions int      `json:"contributions"`
	Languages     []string `json:"languages"`
}

type Error struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}

type Limit struct {
	Message       string `json:"message"`
	Documentation string `json:"documentation_url"`
}
