package utils

import (
	"cli/models"
	"fmt"
	"strings"
)

// MAIN FUNCTION ====================================================================
func DetectLanguages(tree models.TreeNode) ([]string, []string) {

	allFilePaths := TraverseTreeForPaths(tree)
	fmt.Println(allFilePaths)
	goFiles, pyFiles := SeparateFileTypes(allFilePaths)
	return goFiles, pyFiles
}

// TRAVERSE TREE FOR PATHS ==========================================================
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

// SEPERATE FILE TYPES ===============================================================
func SeparateFileTypes(paths []string) ([]string, []string) {
	var goFilePaths = []string{}
	var pyFilePaths = []string{}

	for _, name := range paths {
		switch true {
		case strings.HasSuffix(name, ".go"):
			goFilePaths = append(goFilePaths, name)
		case strings.HasSuffix(name, ".py"):
			pyFilePaths = append(pyFilePaths, name)
		}
	}
	return goFilePaths, pyFilePaths
}
