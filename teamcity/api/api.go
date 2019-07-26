package api

import (
	"encoding/xml"
	"example.com/banana/teamcity/internal/logger"
	"example.com/banana/teamcity/model"
	"io/ioutil"
	"net/http"
	"os"
)

type ITeamcityClient interface {
	FetchProjects() *model.Projects
	FetchBuildTypes() *model.BuildTypes
}

type TeamcityClient struct {}


func (client *TeamcityClient) FetchProjects() *model.Projects {
	var projects model.Projects
	fetchAndUnmarshal("build.innovid.com/httpAuth/app/rest/projects",&projects)
	return &projects
}

func (client *TeamcityClient) FetchBuildTypes() *model.BuildTypes {
	var buildTypes model.BuildTypes
	fetchAndUnmarshal("build.innovid.com/httpAuth/app/rest/buildTypes",&buildTypes)
	return &buildTypes
}

func fetchAndUnmarshal(route string, model interface{}) {
	usr := os.Getenv("user")
	pass := os.Getenv("pass")
	if usr == "" || pass == "" {
		logger.Log.Panic("wtf man... i need creds")
	}

	_ = xml.Unmarshal(
		fetch("https://"+usr+":"+pass+"@"+route),
		model)
}

func fetch(route string) []byte {
	cacheFileName, file, err := readCacheFile(route)
	if err != nil || notOldEnough(file) {
		res, _ := http.Get(route)
		return cacheResponse(res,cacheFileName)
	} else {
		bytes, _ := ioutil.ReadFile(cacheFileName)
		return bytes
	}
}
