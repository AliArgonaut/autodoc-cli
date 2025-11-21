package models

type ConfigTemplate struct {
	Name        string   `json:"app_name"`
	Developer   string   `json:"created by"`
	Description string   `json:"description"`
	Ignore      []string `json:"ignore"`
}

func NewConfigTemplate(name string) ConfigTemplate
