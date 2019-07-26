package main

import (
	"encoding/xml"
	"example.com/banana/teamcity/api/mocks"
	"example.com/banana/teamcity/internal/logger"
	"example.com/banana/teamcity/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

const BuildTypes = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<buildTypes count="2" href="/httpAuth/app/rest/buildTypes">
    <buildType id="boom" name="XXX da"
               description="XXX and friends" projectName="sdfa"
               projectId="bam2"
               href="/httpAuth/app/rest/buildTypes/id:boom"
               webUrl="https://build.banana.com/bam?da=false"/>
    <buildType id="boom" name="XXX sad"
               description="XXX " projectName="da"
               projectId="bam2"
               href="/httpAuth/app/rest/buildTypes/id:boom"
               webUrl="https://build.banana.com/bam?da=false"/>
</buildTypes>
`

const Projects = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<projects count="5" href="/httpAuth/app/rest/projects">
    <project id="boom" name="&lt;Root project&gt;" description="descr"
             href="/httpAuth/app/rest/projects/id:boom"
             webUrl="https://build.banana.com/banana"/>
    <project id="boom" name="sdfa" parentProjectId="_Root"
             href="/httpAuth/app/rest/projects/id:boom"
             webUrl="https://build.banana.com/banana"/>
    <project id="boom" name="Admin" parentProjectId="_Root" description="banana dmin Ap"
             href="/httpAuth/app/rest/projects/id:boom"
             webUrl="https://build.banana.com/banana"/>
    <project id="boom" name="dasdf" parentProjectId="Admin"
             href="/httpAuth/app/rest/projects/id:boom"
             webUrl="https://build.banana.com/banana"/>
    <project id="boom" name="banana" parentProjectId="_Root" description="banana"
             href="/httpAuth/app/rest/projects/id:boom"
             webUrl="https://build.banana.com/banana"/>
</projects>
`

func TestMain(m *testing.M) {
	logger.Log.SetOutput(os.Stdout)
	code := m.Run()
	_ = os.Remove("app.log")
	os.Exit(code)
}

func TestFetchItemsAndFilter(t *testing.T) {

	var buildTypes model.BuildTypes
	var projects model.Projects

	mockTcClient := new(mocks.ITeamcityClient)
	_ = xml.Unmarshal([]byte(BuildTypes), &buildTypes)
	_ = xml.Unmarshal([]byte(Projects), &projects)

	mockTcClient.On("FetchProjects", mock.Anything).Return(&projects)
	mockTcClient.On("FetchBuildTypes", mock.Anything).Return(&buildTypes)

	SearchToken = "XXX"
	filter := fetchItemsAndFilter(mockTcClient)
	mockTcClient.AssertExpectations(t)
	assert.Equal(t, 2, len(filter.Items))

}
