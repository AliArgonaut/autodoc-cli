package models

type TreeNode struct {
	Name     string     `json:"title"`
	Path     string     `json:"path"`
	Children []TreeNode `json:"children"`
}
