package commands

import (
	"encoding/json"
	"example.com/banana/jira/common"
	"fmt"
	"os"
	"time"
)

func  getAactiveSprintId() int {
	return GetAactiveSprint().Id
}


type Sprint struct {
	Id int `json:id`
	Self string `json:self`
	State string `json:state`
	Name string `json:name`
	StartDate time.Time `json:startDate`
	EndDate time.Time `json:endDate`
	OriginBoardId int `json:originBoardId`
	Goal string `json:goal`
}

func (s Sprint) Print(){
	fmt.Println(s.Name)
	fmt.Println("State",s.State)
	fmt.Println(int(time.Until(s.EndDate).Hours()/24),"days left")
	fmt.Println(int(time.Since(s.StartDate).Hours()/24),"days since start date")
}

func  GetAactiveSprint() Sprint {
	request := common.BuilGetRequest("/rest/agile/1.0/board/17/sprint?state=active")
	response := common.Execute(request)
	bytes, _ := json.Marshal(common.JPathGet(response, "$.values[0]"))
	return unmarshal(bytes)
}

func unmarshal(response []byte) Sprint {
	var sprint = Sprint{}
	err := json.Unmarshal(response, &sprint)
	if err != nil {
		println("failed to parse response to Sprint", response)
		os.Exit(1)
	}
	return sprint
}

