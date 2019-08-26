package main

import (
	"bufio"
	"encoding/json"
	"example.com/banana/teamcity/model"
	"fmt"
	"os"
	"strings"
)


type Repo struct {
	name string
	url string
}

func makeAlfredItems(jobs []Repo) model.Items {
	items := model.Items{}
	for idx, job := range jobs {
		items.Items = append(items.Items,
			&model.Item{
				Title:    job.name,
				Arg:      job.url,
				Id:       string(idx)})
	}
	return items
}

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	var repos []Repo
	for scanner.Scan() {
		url := scanner.Text()
		sl := strings.Split(url, "/")
		name := sl[len(sl)-1]
		repos = append(repos, Repo{url: url,name:name})
	}
	bytes, _ := json.Marshal(makeAlfredItems(repos))
	fmt.Println(string(bytes))
}
