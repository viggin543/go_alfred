package model

import "encoding/xml"

type Project struct {
	XMLName xml.Name `xml:"project"`
	Id string `xml:"id,attr"`
	Description string `xml:"description,attr"`
	WebUrl string `xml:"webUrl,attr"`
	Name string `xml:"name,attr"`
}

type Projects struct {
	XMLName xml.Name `xml:"projects"`
	Projects   []Project   `xml:"project"`
}

func  (projects *Projects) ToItems() []*Item {
	items := make([]*Item, len(projects.Projects))
	for i, project := range projects.Projects {
		items[i] = &Item{project.Name, project.WebUrl, project.WebUrl, project.Id}
	}
	return items
}


