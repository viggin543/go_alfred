package main

import (
	"encoding/json"
	"example.com/banana/alfred"
	"example.com/banana/teamcity/api"
	"example.com/banana/teamcity/internal/logger"

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

func fetchItemsAndFilter(client api.ITeamcityClient) alfred.Items {
	logger.Log.Println("searching for teamcity projects/builds with with search token: " + SearchToken)

	var projects = make(chan []*alfred.Item)
	var buildTypes = make(chan []*alfred.Item)
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

func printItemsAsJson(items alfred.Items) {
	itemsJson, _ := json.Marshal(items)
	fmt.Println(string(itemsJson))
}

func filterItemsByCommandLineArg(projects chan []*alfred.Item, buildTypes chan []*alfred.Item) alfred.Items {
	return alfred.Items{
		Items: alfred.FilterItems(
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
