package output

import (
	"encoding/json"
	"example.com/banana/alfred"
	"example.com/banana/github-pull-requests/config"
	"example.com/banana/github-pull-requests/model"
	"fmt"
)

func AlfredPrintPullRequests(pullRequests []model.PullRequest) {
	if config.ExecutedByAlfred {
		var items []*alfred.Item
		for i, pr := range pullRequests {
			items = append(items, &alfred.Item{Title: pr.Link, Id: string(i), Arg: pr.Link, Subtitle: pr.Title})
		}
		if len(items) == 0 {
			items = append(items,&alfred.Item{Title:"no pull requests today...", Subtitle:"push something first..."})
		}
		bytes, _ := json.Marshal(alfred.Items{Items: items})
		fmt.Println(string(bytes))
	}
}
