package main

import (
	"encoding/json"
	"example.com/banana/teamcity/model"
	"fmt"
	"github.com/buger/jsonparser"
	"net/url"
	"os"
)
//in order to install jsonparser
//i  did the following.
//> go get github.com/buger/jsonparser
// go figured this project uses go modules
// and updated the go.mod file

import (
	"io/ioutil"
	"net/http"
)

type Job struct {
	name   string
	url    string
	status string
}

func main() {

	jobs := fetchJobs()
	items := makeAlfredItems(jobs)
	query := getAlfredQuery()
	itemsJson, _ := json.Marshal(
		model.Items{Items: items.FilterByString(query)})
	fmt.Println(string(itemsJson))

}

func getAlfredQuery() string {
	if len(os.Args) == 2 {
		return os.Args[1]
	} else {
		return ""
	}
}

func makeAlfredItems(jobs []Job) model.Items {
	items := model.Items{}
	for idx, job := range jobs {
		items.Items = append(items.Items,
			&model.Item{
				Title:    job.name,
				Subtitle: job.status,
				Arg:      job.url,
				Id:       string(idx)})
	}
	return items
}

func fetchJobs()  []Job {
	jenkinsCreds := os.Getenv("JENKINS_CREDS")
	jenkinsHostPort := os.Getenv("JENKINS_HOST_PORT")
	resp, _ := http.Get(
		fmt.Sprintf(
			"http://%s@%s/api/json/?tree=",
			jenkinsCreds,
			jenkinsHostPort) + url.PathEscape("jobs[name,url,color]"))
	body, _ := ioutil.ReadAll(resp.Body)
	value, _, _, _ := jsonparser.Get(body, "jobs")
	var jobs []Job
	_, _ = jsonparser.ArrayEach(value, func(v []byte, dataType jsonparser.ValueType, offset int, err error) {
		name, _, _, _ := jsonparser.Get(v, "name")
		jobUrl, _, _, _ := jsonparser.Get(v, "url")
		status, _, _, _ := jsonparser.Get(v, "color")
		jobs = append(jobs, Job{name: string(name), url: string(jobUrl), status: string(status)})
	})
	return jobs
}
