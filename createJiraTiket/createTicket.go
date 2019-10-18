package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {


	getEnvValOrPanic("JIRA_USER")
	getEnvValOrPanic("JIRA_PASS")
	getEnvValOrPanic("JIRA_DOMAIN")

	title := flag.String("t", "", "ticket title")
	assignee := flag.String("a", "", "ticket assignee")
	desc := flag.String("d", "", "ticket description")
	flag.Parse()

	if *title == "" || *assignee  == "" || *desc == "" {
		panic("usage: createTicket -t <title> -d <description> -a <assignee>")
	}



}

func getEnvValOrPanic(key string) string {
	user := os.Getenv(key)
	if user == "" {
		fmt.Println("JIRA_USER","JIRA_PASS","JIRA_DOMAIN","plz set env vars")
		panic(fmt.Sprintf("MISSING %s env var", key))
	}
	return user
}
