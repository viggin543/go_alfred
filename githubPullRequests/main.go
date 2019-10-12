package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"os"
)

var pathToLocalRepositories string
var githubUser string
var githubPass string

func init() {
	githubUser = getEnvVar("GITHUB_USER")
	githubPass = getEnvVar("GITHUB_PASS")
	pathToLocalRepositories = getEnvVar("REPOS_DIR")
}

func getEnvVar(s string) string {
	if os.Getenv(s) != "" {
		return  os.Getenv(s)
	} else {
		panic("please set REPOS_DIR env var")
	}
}

func main() {
	pullRequests := parallelGetPullRequests(localRepos())
	fmt.Println(pullRequests)
}

func parallelGetPullRequests(repos []string) []PullRequest {
	pullRequests := make(chan []PullRequest)
	for _, repo := range repos {
		go func(repo string) {
			requests := getPullRequests(repo)
			pullRequests <- requests
		}(repo)
	}
	var prr []PullRequest
	for range repos {
		requests := <-pullRequests
		for _, pr := range requests {
			prr = append(prr, pr)
		}
	}
	return prr
}

func localRepos() []string {
	files, _ := ioutil.ReadDir(pathToLocalRepositories)
	var repos []string
	for _, file := range files {
		repos = append(repos, file.Name())
	}
	return repos
}


type PullRequest struct {
	user string
	link string
}

func getPullRequests(repo string) []PullRequest {
	url := pullRequestUrl(repo)
	body := fetchPullRequests(url)
	return marshal(body)
}

func marshal(body *[]byte) []PullRequest {
	var prs []PullRequest
	_, _ = jsonparser.ArrayEach(*body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		url, _, _, _ := jsonparser.Get(value, "html_url")
		user, _, _, _ := jsonparser.Get(value, "user")
		userLogin, _, _, _ := jsonparser.Get(user, "login")
		if len(userLogin) + len(url) > 0 {
			prs = append(prs, PullRequest{ string(userLogin),string(url)})
		}
	})
	return prs
}

func fetchPullRequests(url string) *[]byte {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return &body
}

func pullRequestUrl(repo string) string {
	return fmt.Sprintf("https://%s:%s@api.github.com/repos/tg-17/%s/pulls?state=open",
		 githubUser, githubPass,repo)
}
