package models

type ConfigTemplate struct {
	name      `json:"app_name"`
	developer `json: "created by"`
	ignore    []string
}
