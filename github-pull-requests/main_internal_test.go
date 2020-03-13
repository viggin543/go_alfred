package main

import (
	"example.com/banana/github-pull-requests/api"
	"example.com/banana/github-pull-requests/config"
	"example.com/banana/github-pull-requests/model"
	"fmt"
	"strings"
	"testing"
)


func TestGetPullRequests(t *testing.T) {
	pr := api.GetPullRequests("svc-subject")
	if len(pr) != 1 {
		fmt.Println("len(pr)", len(pr))
		t.Fail()
	}
	if !strings.Contains(pr[0].Link, "http") {
		fmt.Println(pr[0])
		fmt.Println("fail invalid link")
		t.Fail()
	}
	if pr[0].User == "" {
		fmt.Println(pr[0])
		fmt.Println("fail invalid user")
		t.Fail()
	}
}

func TestParallelGetPrs(t *testing.T) {
	requests := parallelGetPullRequests(config.LocalRepos())
	fmt.Println("got pull requests", requests)
	if len(requests) == 0 {
		t.Fail()
	}
}

func TestFilterPrs(t *testing.T) {
	config.GithubUser = "me"
	var pullRequests = []model.PullRequest{{User: "me",Link:"linkkk"},{User: "me_not",Link:"linkkk"}}

	pullRequests = filterMyPullRequests(pullRequests)
	if len(pullRequests) == 0 {
		t.Log(pullRequests,"expected one result")
		t.Fail()
	}
	if pullRequests[0].User != "me" {
		t.Log("should filter me",pullRequests[0].User)
		t.Fail()
	}

}
