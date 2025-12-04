package models

type AgentRequestParams struct {
	AppName     string     `json:"app_name"`
	Description string     `json:"description"`
	Filetree    TreeNode   `json:"Filetree"`
	Contents    []CodeFile `json:"Code"`
}
