package main

import (
	"example.com/banana/github-pull-requests/api"
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

