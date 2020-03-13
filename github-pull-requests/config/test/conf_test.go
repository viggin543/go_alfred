package test

import (
	"example.com/banana/github-pull-requests/config"
	"testing"
)

func TestGetRepos(t *testing.T) {
	repos := config.LocalRepos()
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

