package main

import (
	"example.com/banana/jira/commands"
	"flag"
	"fmt"
	"os"
)


var command = flag.String("c", "ls", "command")


func init() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n ", os.Args[0])
		fmt.Printf("Usage: jira -t <title> -d <description> -a <assignee> -c create|ct")
		fmt.Printf("Usage: jira  -c list|ls")
		flag.PrintDefaults()
	}
}

func main() {

	flag.Parse()

	switch *command {
	case "list", "ls":
		commands.NewListProjectTeamCommand().
			Execute()
	case "create", "ct":
		commands.NewCreateIssue().
			ParseFlags().
			Execute()
	default:
		flag.Usage()
	}

}



