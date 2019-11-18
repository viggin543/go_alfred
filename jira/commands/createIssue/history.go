package createIssue

import "example.com/banana/jira/common"

var history_file = "~/.jira_tickets"
var epics_file = "~/.jira_epics"

func PrintCreatedTasks(){
	common.PrintFileContent(history_file)
}

func PrintEpics(){
	common.PrintFileContent(epics_file)
}
