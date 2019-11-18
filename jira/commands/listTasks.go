package commands

import (
	"bytes"
	"example.com/banana/jira/common"
	"flag"
	"fmt"
)

var verbose = flag.Bool("v", false, "verbose output")

type listIssues struct {
	verbose bool
}

func NewListIssuesCommand() *listIssues {
	return &listIssues{verbose: false}
}

func (l *listIssues) ParseArgs() *listIssues {
	flag.Parse()
	l.verbose = *verbose
	return l
}

func (l *listIssues) Execute() {
	l.print(GetIssues())
}

func GetIssues() *[]interface{} {
	path := "/rest/api/2/search"
	body := fmt.Sprintf(`{"jql":
									"Assignee = \"%s\" AND Status in (\"To Do\" ,\"In Progress\",\"Review\")",
									"startAt":0,
									"maxResults":30}`, user)
	req := common.BuildPostRequest(
		path,
		bytes.NewBuffer([]byte(body)))
	resp := common.Execute(req)
	issues := common.JPathGet(resp, "$.issues").([]interface{})
	return &issues
}

func (l *listIssues) print(issues *[]interface{}) {
	if l.verbose {
		fmt.Println("=========================================================================================")
		for _, v := range *issues {
			issue := v.(map[string]interface{})
			key := issue["key"]
			fmt.Println(key)
			summary := issue["fields"].(map[string]interface{})["summary"]
			fmt.Println(summary)
			fmt.Println(fmt.Sprintf("https://tg17home.atlassian.net/browse/%s", key))
			fmt.Println("=========================================================================================")
		}
	} else {
		for _, v := range *issues {
			ticket := v.(map[string]interface{})["key"]
			fmt.Println(ticket)
		}
	}
}
