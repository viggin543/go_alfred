package main

import (
	"encoding/json"
	"example.com/banana/alfred"
	"flag"
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"os"
)

var pathToLocalRepositories string
var githubUser string
var githubPass string
var executedByAlfred bool
var org string

func init() {
	isAlfred := flag.Bool("alfred", false, "a boolean value ")
	flag.Parse()
	executedByAlfred = *isAlfred
	githubUser = getEnvVarOrPanic("GITHUB_USER")
	githubPass = getEnvVarOrPanic("GITHUB_PASS")
	org = getEnvVarOrPanic("GITHUB_ORG")
	pathToLocalRepositories = getEnvVarOrPanic("REPOS_DIR")
}

func getEnvVarOrPanic(s string) string {
	if os.Getenv(s) != "" {
		return  os.Getenv(s)
	} else {
		panic("please set REPOS_DIR env var")
	}
}

func main() {
	pullRequests := parallelGetPullRequests(localRepos())
	prettyPrintPullRequests(pullRequests)
	alfredPrintPullRequests(pullRequests)
}

func alfredPrintPullRequests(pullRequests []PullRequest) {
	if executedByAlfred {
		var items []*alfred.Item
		for i, pr := range pullRequests {
			if pr.user == githubUser {
				items = append(items, &alfred.Item{Title: pr.link, Id: string(i),Arg:pr.link,Subtitle:pr.user})
			}
		}
		if len(items) == 0 {
			items = append(items,&alfred.Item{Title:"no pull requests today...", Subtitle:"push something first..."})
		}
		bytes, _ := json.Marshal(alfred.Items{Items: items})
		fmt.Println(string(bytes))
	}
}

func prettyPrintPullRequests(pullRequests []PullRequest) {
	if !executedByAlfred {
		fmt.Println("Your pull requests:")
		fmt.Println("")
		if len(pullRequests) == 0 {
			fmt.Println("Sorry.. it seems you have no open pull requests")
		}
		for _, pr := range pullRequests {
			if pr.user == githubUser {
				fmt.Println(pr.link)
			}
		}
	}
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
	return fmt.Sprintf("https://%s:%s@api.github.com/repos/%s/%s/pulls?state=open",
		 githubUser, githubPass,org,repo)
}
