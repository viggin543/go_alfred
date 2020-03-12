package main

import (
	"fmt"
	"strings"
	"testing"
)


func TestGoogleRequest(t *testing.T) {
	pr := getPullRequests("svc-subject")
	if len(pr) != 1 {
		fmt.Println("len(pr)", len(pr))
		t.Fail()
	}
	if !strings.Contains(pr[0].link, "http") {
		fmt.Println(pr[0])
		t.Fail()
	}
	if pr[0].user != "viggin543" {
		fmt.Println(pr[0])
		t.Fail()
	}
}

func TestGetRepos(t *testing.T) {
	repos := localRepos()
	if len(repos) == 0 {
		t.Fail()
	}
	if !contains(repos, "svc-web") {
		t.Fail()
	}
}

func contains(highstack []string, niddle string) bool {
	pass := false
	for _, elm := range highstack {
		if elm == niddle {
			pass = true
		}
	}
	return pass
}
