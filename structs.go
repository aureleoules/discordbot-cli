package main

//Repository : GitHub Repository Object
type Repository struct {
	HTMLURL string `json:"html_url"`
}

//GitHubAPIResponse : []Repository
type GitHubAPIResponse struct {
	Repositories []Repository `json:"items"`
}

//Config : JSON Config file
type Config struct {
	GitHubToken string `json:"GitHubToken"`
}
