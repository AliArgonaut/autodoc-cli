package utils

import (
	"cli/models"
)

func TraverseTreeForPaths(tree models.TreeNode) []string {
	allPaths := []string{}
	if tree.FileType == "f" {
		allPaths = append(allPaths, tree.Path)
	} else {
		for i, _ := range tree.Children {
			childPaths := TraverseTreeForPaths(tree.Children[i])
			allPaths = append(allPaths, childPaths...)
		}
	}
	return allPaths
}
