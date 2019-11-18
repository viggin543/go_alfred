package main

import (
	"example.com/banana/jira/commands"
	"example.com/banana/jira/commands/createIssue"
	"flag"
	"fmt"
	"os"
)


var command = flag.String("c", "ls", "command")


func init() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n ", os.Args[0])
		fmt.Printf("Usage: jira -c create|ct -t <title> -d <description> -a <assignee> [-e <epic task id>]")
		fmt.Printf("creates a task, for the current sprint, the assignee is searched in the team members ist")
		fmt.Printf("Usage: jira  -c list|ls")
		fmt.Printf("Usage: jira  -c sprint")
		flag.PrintDefaults()
	}
}

func main() {

	flag.Parse()

	switch *command {
	case "team":
		commands.NewListProjectTeamCommand().Execute()
	case "create", "ct":
		createIssue.NewCreateIssue().
			ParseFlags().
			Execute()
	case "sprint":
		commands.GetAactiveSprint().
			Print()
	case "list","ls":
		commands.NewListIssuesCommand().
			ParseArgs().
			Execute()
	case "history":
		createIssue.PrintCreatedTasks()
	case "epics":
		createIssue.PrintEpics()
	default:
		flag.Usage()
	}

}



