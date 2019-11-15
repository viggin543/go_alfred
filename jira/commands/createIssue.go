package commands

import (
	"bytes"
	"example.com/banana/jira/common"
	"flag"
	"fmt"
	"os"
)

var title = flag.String("t", "", "ticket title")
var assignee = flag.String("a", "", "ticket assignee")
var desc = flag.String("d", "", "ticket description")
var epic = flag.Int("e", 0, "epic ticket")
var _, _, domain = common.Config()

func NewCreateIssue() createIssue {
	return createIssue{}
}

//createIssue ...
type createIssue struct {
	title       string
	assignee    string
	description string
}

//ParseFlags parse flags title assignee description
// jira -t <title> -d <description> -a <assignee> create
func (t createIssue) ParseFlags() createIssue {
	flag.Parse()
	t.title = *title
	t.assignee = *assignee
	t.description = *desc
	t.assert()
	return t
}

//Execute create jira issue
func (t createIssue) Execute() interface{} {
	req := common.BuildPostRequest("/rest/api/2/issue/", t.postBody())
	body := common.Execute(req)
	taskNumber := common.ParseToSting(body, "$.key")
	createdTask := fmt.Sprintf("https://%s/browse/%s", domain, taskNumber)
	common.AppendToFile("~/.jira_tickets",createdTask)
	fmt.Println(createdTask)
	return nil
}

func (t createIssue) getAactiveSprintId() int {
	r := common.BuilGetdRequest("/rest/agile/1.0/board/17/sprint?state=active")
	b := common.Execute(r)
	sprintId := common.ParseToSplitInt(b, "$.values..id")[0]
	return sprintId
}

func (t createIssue) postBody() *bytes.Buffer {

	body := bytes.NewBuffer([]byte(fmt.Sprintf(`{
	"fields": {
	   "project": {"key": "UD"},
	   "summary": "%s",
	   "description": "%s",
	   "customfield_10064": {"value": "Backend"},
		"customfield_10010":%d,
	   "issuetype": {"name": "Task"},
	   "assignee": {"name":"%s"}
		%s
		}
	}`,
		t.title,
		t.description,
		t.getAactiveSprintId(),
		t.assignee,
		getEpicLink())))
	return body
}

func getEpicLink() string {
	if *epic != 0 {
		return fmt.Sprintf(`,"customfield_10008":"UD-%d"`, *epic)
	} else {
		return ""
	}
}

func (t createIssue) assert() {
	if t.assignee == "" || t.description == "" || t.title == "" {
		flag.Usage()
		os.Exit(1)
	}
}
