package commands

import (
	"example.com/banana/jira/common"
	"fmt"
)




func NewListProjectTeamCommand() listProjectTeam {
	return listProjectTeam{quiet:false}
}

//listProjectTeam ...
type listProjectTeam struct {
	quiet bool
}

func (t listProjectTeam) NoLogs() listProjectTeam {
	t.quiet = true
	return t
}

//Execute ...
func (t listProjectTeam) Execute() []string {
	req := common.BuilGetdRequest("/rest/api/2/user/assignable/search?project=UD")
	body := common.Execute(req)
	jiraUsers := common.ParseToSplitStr(body, "$..name")
	t.print(jiraUsers)
	return jiraUsers
}

func (t listProjectTeam) print(jiraUsers []string) {
	if !t.quiet {
		for _, user := range jiraUsers {
			fmt.Println(user)
		}
	}

}
