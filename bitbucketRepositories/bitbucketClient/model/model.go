package model


type Html struct {
	Href string `json:"href"`
}
type Link struct {
	Html Html `json:"html"`
}
type Repository struct {
	Name  string `json:"name"`
	Links Link   `json:"links"`
}

type RepositoriesResponse struct {
	Values []Repository `json:"values"`
}