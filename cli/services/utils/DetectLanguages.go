package utils

import (
	"cli/models"
	"fmt"
	"strings"
)

func DetectLanguages(tree models.TreeNode) ([]string, []string) {

	allFilePaths := TraverseTreeForPaths(tree)
	fmt.Println("all file paths vvvvv")
	fmt.Println(allFilePaths)
	goFiles, pyFiles := SeparateFileTypes(allFilePaths)
	return goFiles, pyFiles
}

func TraverseTreeForPaths(tree models.TreeNode) []string {
	allPaths := []string{}
	if tree.FileType == "File" {
		fmt.Println("filetype is file path::::::::::", tree.Path)
		allPaths = append(allPaths, tree.Path)
		return allPaths
	} else {
		for i, _ := range tree.Children {
			TraverseTreeForPaths(tree.Children[i])
		}
	}
	fmt.Println("all PATHS INNER FUNC:=====", allPaths)
	return allPaths
}

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
