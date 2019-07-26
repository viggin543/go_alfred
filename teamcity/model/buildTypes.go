package model


import "encoding/xml"

type BuildType struct {
	XMLName xml.Name `xml:"buildType"`
	Id string `xml:"id,attr"`
	Description string `xml:"description,attr"`
	WebUrl string `xml:"webUrl,attr"`
	Name string `xml:"name,attr"`
}

type BuildTypes struct {
	XMLName    xml.Name    `xml:"buildTypes"`
	BuildTypes []BuildType `xml:"buildType"`
}

func  (buildTypes *BuildTypes) ToItems() []*Item {
	items := make([]*Item, len(buildTypes.BuildTypes))
	for i, project := range buildTypes.BuildTypes {
		items[i] = &Item{project.Name, project.WebUrl, project.WebUrl, project.Id}
	}
	return items
}



