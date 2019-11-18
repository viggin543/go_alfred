package createIssue

import (
	"bytes"
	"example.com/banana/jira/commands"
	"example.com/banana/jira/common"
	"flag"
	"fmt"
	"os"
	"strings"
)

var Title = flag.String("t", "", "ticket Title")
var Assignee = flag.String("a", "", "ticket Assignee")
var Desc = flag.String("d", "", "ticket Description")
var Epic = flag.Int("e", 0, "Epic ticket")

var _, _, domain = common.Config()


func NewCreateIssue() *createIssue {
	return &createIssue{}
}

type createIssue struct {
	Title       string
	Assignee    string
	Description string
	Epic int
}


func (t *createIssue) ParseFlags() *createIssue {
	flag.Parse()
	t.Title = *Title
	t.Description = *Desc
	t.parseEpic()
	t.setAssignee().
		assert()
	return t
}

func (t *createIssue) parseEpic() {
	if *Epic != 0 {
		common.AppendToFile(epics_file, string(*Epic))
	}
	t.Epic = *Epic
}

func (t *createIssue) setAssignee() *createIssue {
	team := commands.NewListProjectTeamCommand().NoLogs().Execute()
	for _, member := range team {
		if strings.Contains(member, *Assignee) {
			t.Assignee = member
			break
		}
	}
	if t.Assignee == "" && *Assignee != "" {
		fmt.Println("cant fine assignee",*Assignee)
		os.Exit(1)
	}
	return t
}

func (t *createIssue) Execute()  {
	req := common.BuildPostRequest("/rest/api/2/issue/", t.postBody())
	body := common.Execute(req)
	taskNumber := common.ParseToSting(body, "$.key")
	createdTask := fmt.Sprintf("https://%s/browse/%s", domain, taskNumber)

	common.AppendToFile(history_file,createdTask)
	fmt.Println(createdTask)
}


func (t *createIssue) postBody() *bytes.Buffer {

	body := bytes.NewBuffer([]byte(fmt.Sprintf(`{
	"fields": {
	   "project": {"key": "UD"},
	   "summary": "%s",
	   "Description": "%s",
	   "customfield_10064": {"value": "Backend"},
		"customfield_10010":%d,
	   "issuetype": {"name": "Task"},
	   "Assignee": {"name":"%s"}
		%s
		}
	}`,
		t.Title,
		t.Description,
		commands.GetAactiveSprint().Id,
		t.Assignee,
		t.getEpicLink())))
	return body
}

func (t *createIssue) getEpicLink() string {
	if t.Epic != 0 {
		return fmt.Sprintf(`,"customfield_10008":"UD-%d"`, t.Epic)
	} else {
		return ""
	}
}

func (t *createIssue) assert() {
	if t.Assignee == "" || t.Description == "" || t.Title == "" {
		flag.Usage()
		os.Exit(1)
	}
}
