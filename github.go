package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//GetRepositories : Fetch repositories on GitHub
func GetRepositories(page string, token string) []Repository {
	reqURL := APIEndpoint + "/search/code"
	client := &http.Client{}

	//Prepare request
	req, err := http.NewRequest(
		"GET",
		reqURL,
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	//Pass data to request (params & headers)
	params := req.URL.Query()
	params.Add("q", ".login( discord+language:javascript")
	params.Add("sort", "indexed")
	params.Add("order", "desc")
	params.Add("page", page)
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Authorization", "token "+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//Make request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	//Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error: ", err)
		return nil
	}

	//Parse body into Repositories
	var APIReponse GitHubAPIResponse
	err = json.Unmarshal([]byte(body), &APIReponse)
	if err != nil {
		log.Fatal(err)
	}

	return APIReponse.Repositories
}

//DownloadCode : Download code of a list of file from GitHub
func DownloadCode(repositories []Repository) []string {
	//Process URL in order to get the direct download link of the JS file
	processURLs := func(repositories []Repository) []string {
		var list []string
		for i := range repositories {
			url := repositories[i].HTMLURL
			url = strings.Replace(repositories[i].HTMLURL, "https://github.com/", "https://raw.githubusercontent.com/", 1)
			url = strings.Replace(url, "blob/", "", -1)
			list = append(list, url)
		}
		return list
	}

	URLs := processURLs(repositories)

	var codeList []string

	for i := range URLs {
		url := URLs[i] //Set URL to current

		//Request file
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		//Parse file
		code := string(body)

		//Append to list
		codeList = append(codeList, code)
	}

	return codeList
}

//AnalyseCode : Look for tokens in a list of code
func AnalyseCode(codeList []string) []string {

	getStringInBetween := func(str string, start string, end string) (result string) {
		s := strings.Index(str, start)
		if s == -1 {
			return
		}
		s += len(start)
		e := strings.Index(str, end)
		return str[s:e]
	}

	var tokens []string

	for i := range codeList {
		code := codeList[i]
		lines := strings.Split(code, "\n")

		for j := range lines {
			line := lines[j]

			if strings.Contains(line, ".login('") || strings.Contains(line, ".login(\"") || strings.Contains(line, ".login(`") {
				token := getStringInBetween(line, "('", "')")
				if token != "" {
					tokens = append(tokens, token)
				}
			}
		}
	}
	return tokens
}
