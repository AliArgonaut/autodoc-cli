package utils

import (
	"cli/models"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func BuildTree(cwd string, name string, ignored []string) models.TreeNode {
	Treenode := models.TreeNode{}
	Treenode.Name = name
	Treenode.Path = cwd

	files, err := os.ReadDir(cwd)
	if err != nil {
		log.Fatalln("couldnt read dir")
	}

	for _, file := range files {
		isIgnored := false

		for _, ignore := range ignored {
			if strings.Contains(file.Name(), ignore) {
				isIgnored = true
			}
		}

		if !isIgnored && file.IsDir() {
			Treenode.Children = append(Treenode.Children, BuildTree(filepath.Join(cwd, file.Name()), file.Name(), ignored))
		} else if !isIgnored && !file.IsDir() {
			Treenode.Children = append(Treenode.Children, models.TreeNode{
				Name: file.Name(),
				Path: filepath.Join(cwd, file.Name()),
			})
		} else {
			continue
		}

	}
	return Treenode
}

//	entries, err := os.ReadDir(cwd)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	for _, entry := range entries {
//		if entry.IsDir() {
//			Treenode.Children = append(Treenode.Children, BuildTree(filepath.Join(cwd, entry.Name()), entry.Name()))
//		} else {
//			Treenode.Children = append(Treenode.Children, models.TreeNode{
//				Name: entry.Name(),
//				Path: filepath.Join(cwd, entry.Name()),
//			})
//		}
//	}
//
//	return Treenode
