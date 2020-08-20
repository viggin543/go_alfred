package bitbucketClient

import (
	"encoding/json"
	"example.com/banana/bitbucketRepositories/bitbucketClient/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Repo struct {
	Name, Url string
}


var pagesCount = 10

func GetRepositories() []Repo {
	getResultsChan := parallelGetRepositories()
	var rr []Repo

	for element := range getResultsChan{
		for _, v := range element.Values {
			rr = append(rr, Repo{Name: v.Name, Url: v.Links.Html.Href})
		}
	}

	return rr
}

type asyncReq struct {
	waitGroup *sync.WaitGroup
	page           int
	channel        chan model.RepositoriesResponse
}

func parallelGetRepositories() chan model.RepositoriesResponse {
	getResultsChan := make(chan model.RepositoriesResponse, pagesCount)
	var waitGroup sync.WaitGroup
	waitGroup.Add(pagesCount)
	req := asyncReq{
		waitGroup: &waitGroup,
		channel:   getResultsChan,
	}
	for page := 1; page <= pagesCount; page++ {
		req.getAsync(page)
	}
	waitGroup.Wait()
	close(getResultsChan)
	return getResultsChan
}

func (r *asyncReq) getAsync(page int) {
	go func(idx int) {
		resp, err := http.Get(makeUrl(idx))
		if err != nil {
			panic(err.Error())
		}
		page := fetchPage(resp)
		if len(page.Values) > 0 {
			r.channel <- page
		}
		r.waitGroup.Done()
	}(page)
}

func fetchPage(resp *http.Response) model.RepositoriesResponse {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("failed to read resp")
	}
	var response model.RepositoriesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic("failed to unmarshal")
	}
	return response
}

func makeUrl(page int) string {
	return fmt.Sprintf("https://%s:%s@api.bitbucket.org/2.0/repositories/adikatech?page=%d", user, pass, page)
}
