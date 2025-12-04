package models

type TreeNode struct {
	Name     string     `json:"name"`
	Path     string     `json:"-"`
	FileType string     `json:"FileType"`
	Children []TreeNode `json:"contains,omitempty"`
}

//important to give this context to the agents layer that n is name, p is path, etc etc etc..it can be in teh agent instruction.
