package main

import (
	"fmt"

	"example.com/banana/jira/common"
)

type Project struct {
	Key string `json:"key"`
}

type Team struct {
	Value string `json:"value"`
}

type IssueType struct {
	Name string `json:"name"`
}

type Assignee struct {
	Name string `json:"name"`
}

type Fields struct {
	Project     Project   `json:"project"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Team        Team      `json:"customfield_10064"`
	IssueType   IssueType `json:"issuetype"`
	Assignee    Assignee  `json:"assignee"`
}

type Ticket struct {
	Fields Fields `json:"fields"`
}

type CreateTicketResp struct {
	id   string
	key  string
	self string
}

func (t *CreateTicketResp) publicUrl() string {
	_, _, domain := common.Config()
	return fmt.Sprintf("https://%s/browse/%s", domain, t.key)
}

func makeTicket(assignee *string, desc *string, title *string) Ticket {
	return Ticket{Fields: Fields{
		Assignee:    Assignee{Name: *assignee},
		Project:     Project{Key: "UD"},
		Description: *desc,
		Summary:     *title,
		Team:        Team{Value: "Backend"},
		IssueType:   IssueType{Name: "Task"},
	}}
}
