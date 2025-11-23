package models

type TreeNode struct {
	Name     string     `json:"title"`
	Path     string     `json:"path"`
	FileType string     `json:"FileType"`
	Children []TreeNode `json:"children"`
}
