package models

type Config struct {
	Name        string   `json:"name"`
	Developer   string   `json:"developer"`
	Description string   `json:"description"`
	Ignore      []string `json:"ignore"`
}
