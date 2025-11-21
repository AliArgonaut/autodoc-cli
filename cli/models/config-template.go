package models

type ConfigTemplate struct {
	name        string   `json:"app_name"`
	developer   string   `json:"created by"`
	description string   `json:description`
	ignore      []string `json:"ignore"`
}
