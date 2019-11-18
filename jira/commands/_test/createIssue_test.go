package _test

import (
	"example.com/banana/jira/commands/createIssue"
	"os"
	"os/exec"
	"testing"
)

var _title = "title"
var _ass = "igor"
var _desc = "desc"

func TestSetAssignee(t *testing.T) {
	createIssue.Title = &_title
	createIssue.Assignee = &_ass
	createIssue.Desc = &_desc
	command := createIssue.NewCreateIssue().ParseFlags()
	if command.Assignee != "igor.domrev" {
		t.Fatal("invalid assignee...expected igor.domrev got",command.Assignee)
	}
}

func TestAssertFlags(t *testing.T) {
	empty := ""
	createIssue.Title = &empty
	createIssue.Assignee = &empty
	createIssue.Desc = &empty
	if os.Getenv("BE_CRASHER") == "1" {
		createIssue.NewCreateIssue().ParseFlags()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestAssertFlags")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}


