package model

import (
	"encoding/xml"
	"example.com/banana/alfred"
)

type Project struct {
	XMLName     xml.Name `xml:"project"`
	Id          string   `xml:"id,attr"`
	Description string   `xml:"description,attr"`
	WebUrl      string   `xml:"webUrl,attr"`
	Name        string   `xml:"name,attr"`
}

type Projects struct {
	XMLName  xml.Name  `xml:"projects"`
	Projects []Project `xml:"project"`
}

func (projects *Projects) ToItems() []*alfred.Item {
	items := make([]*alfred.Item, len(projects.Projects))
	for i, project := range projects.Projects {
		items[i] = &alfred.Item{project.Name, project.WebUrl, project.WebUrl, project.Id}
	}
	return items
}
