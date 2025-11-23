package utils

import "cli/models"

func DetectLanguage(tree models.TreeNode) ([]string, []string) {

	emptyArray := []string{}
	allFilePaths := TraverseTreeForPaths(tree, emptyArray)

	///	goFilePaths := []string{}
	///	pyFilePaths := []string{}

}

func TraverseTreeForPaths(tree models.TreeNode, allFilePaths []string) []string {

	if tree.FileType == "File" {
		allFilePaths = append(allFilePaths, tree.Path)
	} else {
		for i, _ := range tree.Children {
			TraverseTreeForPaths(tree.Children[i], allFilePaths)
		}
	}
	return allFilePaths
}
