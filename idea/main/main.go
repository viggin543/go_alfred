package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	projects := getLatestProjectsPaths(readIdeaRecentProjectsXml())
	var items = &Items{makeAlfredItems(projects)}
	items = filterItems(items)
	bytes, _ := json.Marshal(items)
	fmt.Println(string(bytes))
}

type Item struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Arg      string `json:"arg"`
	Id       int    `json:"id"`
}
type Items struct {
	Items []Item `json:"items"`
}

func (items *Items) filter(query string) *Items {
	var filtered []Item
	for _, item := range items.Items {
		if strings.Contains(item.Title, query) {
			filtered = append(filtered, item)
		}
	}
	return &Items{filtered}
}

func filterItems(items *Items) *Items {
	if len(os.Args) >= 2 {
		query := os.Args[1]
		items = items.filter(query)
	}
	return items
}

func makeAlfredItems(projects []string) []Item {
	var items []Item
	for idx, path := range projects {
		split := strings.Split(path, "/")
		items = append(items, Item{split[len(split)-1], path, path, idx})
	}
	return items
}

func getLatestProjectsPaths(b []byte) []string {
	var projects []string

	for _, line := range strings.Split(strings.TrimSuffix(string(b), "\n"), "\n") {
		if strings.Contains(line, "option value") {
			project := String{line}.inUserHomeDir().noCloseTag().outsideUserHomeDir().withoutWhitespace().Str
			projects = append(projects, project)
		}
	}
	return projects
}

func readIdeaRecentProjectsXml() []byte {
	pathToXml := os.Getenv("IDEA_RECENT_PROJECTS")
	b, _ := ioutil.ReadFile(pathToXml)
	return b
}
