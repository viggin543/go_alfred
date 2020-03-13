package api

import (
	"example.com/banana/github-pull-requests/config"
	"example.com/banana/github-pull-requests/model"
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
)

func GetPullRequests(repo string) []model.PullRequest {
	url := pullRequestUrl(repo)
	body := fetchPullRequests(url)
	return marshal(body)
}

func marshal(body *[]byte) []model.PullRequest {
	var prs []model.PullRequest
	jsonparser.ArrayEach(*body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		url, _, _, _ := jsonparser.Get(value, "html_url")
		user, _, _, _ := jsonparser.Get(value, "user")
		userLogin, _, _, _ := jsonparser.Get(user, "login")
		if len(userLogin) + len(url) > 0 {
			prs = append(prs, model.PullRequest{ string(userLogin),string(url)})
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
		config.GithubUser, config.GithubPass,config.Org,repo)
}
