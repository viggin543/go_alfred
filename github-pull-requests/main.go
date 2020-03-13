package main

import (
	"example.com/banana/github-pull-requests/api"
	"example.com/banana/github-pull-requests/config"
	"example.com/banana/github-pull-requests/model"
	"example.com/banana/github-pull-requests/output"
)


type Prs []model.PullRequest

func main() {
	pullRequests := parallelGetPullRequests(config.LocalRepos()).filterMyPullRequests()
	output.PrettyPrintPullRequests(pullRequests)
	output.AlfredPrintPullRequests(pullRequests)
}



func ( prs Prs) filterMyPullRequests() Prs {
	var res = make([]model.PullRequest,0,len(prs))
	for _,pr := range prs {
		if pr.User == config.GithubUser {
			res = append(res,pr)
		}
	}
	return res
}

func parallelGetPullRequests(repos []string) Prs {
	pullRequests := make(chan []model.PullRequest)
	for _, repo := range repos {
		go func(repo string) {
			requests := api.GetPullRequests(repo)
			pullRequests <- requests
		}(repo)
	}
	var prr []model.PullRequest
	for range repos {
		requests := <-pullRequests
		for _, pr := range requests {
			prr = append(prr, pr)
		}
	}
	return prr
}




