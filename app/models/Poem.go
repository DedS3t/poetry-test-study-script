package models

type PoemDto struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Poem   string `json:"poem"`
}
