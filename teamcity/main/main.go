package main

import (
	"encoding/json"
	"example.com/banana/teamcity/api"
	//logger is an internal package. in can be consumed within teamcity module
	//teamcity => internal
	// cant be consumed outside. very nice
	"example.com/banana/teamcity/internal/logger"
	//https://blog.learngoprogramming.com/special-packages-and-directories-in-go-1d6295690a6b
	"example.com/banana/teamcity/model"
	"fmt"
	"os"
)

var SearchToken string

func main() {

	var client api.ITeamcityClient = &api.TeamcityClient{}

	SearchToken = commandLineArg()

	items := fetchItemsAndFilter(client)

	printItemsAsJson(items)
}

func fetchItemsAndFilter(client api.ITeamcityClient) model.Items {
	logger.Log.Println("searching for teamcity projects/builds with with search token: " + SearchToken)

	var projects = make(chan []*model.Item)
	var buildTypes = make(chan []*model.Item)
	go func() {
		fetchProjects := client.FetchProjects()
		projects <- fetchProjects.ToItems()
	}()
	go func() {
		buildTypes <- client.FetchBuildTypes().ToItems()
	}()
	items := filterItemsByCommandLineArg(projects, buildTypes)
	return items
}

func printItemsAsJson(items model.Items) {
	itemsJson, _ := json.Marshal(items)
	fmt.Println(string(itemsJson))
}

func filterItemsByCommandLineArg(projects chan []*model.Item, buildTypes chan []*model.Item) model.Items {
	return model.Items{
		Items: model.FilterItems(
			append(<-projects, <-buildTypes...),
			SearchToken)}
}

func commandLineArg() string {
	if len(os.Args) == 1 {
		return ""
	} else {
		return os.Args[1]
	}
}
