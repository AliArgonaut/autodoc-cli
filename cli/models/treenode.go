package models

type TreeNode struct {
	Name     string     `json:"n"`
	Path     string     `json:"p"`
	FileType string     `json:"t"`
	Children []TreeNode `json:"c"`
}
