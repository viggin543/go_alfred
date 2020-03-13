package output

import (
	"example.com/banana/github-pull-requests/config"
	"example.com/banana/github-pull-requests/model"
	"fmt"
)

func PrettyPrintPullRequests(pullRequests []model.PullRequest) {
	if !config.ExecutedByAlfred {
		fmt.Println("Your pull requests:")
		fmt.Println("")
		if len(pullRequests) == 0 {
			fmt.Println("Sorry.. it seems you have no open pull requests")
		}
		for _, pr := range pullRequests {
			if pr.User == config.GithubUser {
				fmt.Println(pr.Link)
			}
		}
	}
}
