package services

import (
	"cli/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GenerateService() {
	fmt.Println("generate service hit") //check!

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("failed to get current working directory")
	}

	cwdName := filepath.Base(cwd)

	//ping python backend with post request containing project metadata
}

//helper functions for getting file tree structure and so on
//
//
//
// reads given dir and returns a slice of names. func ReadDir(name string) ([]DirEntry, error)
// has methods like .isDir(), and .Name()

func BuildTree(cwd string, name string) models.TreeNode {
	Treenode := models.TreeNode{}

	Treenode.Name = name
	Treenode.Path = cwd

	recursiveCD, err := os.Getwd()

	entries, err := os.ReadDir(cwd)
	if err != nil {
		log.Fatalln("got cwd but failed to read entries")
	}
	for _, entry := range entries {
		if entry.IsDir() {
			Treenode.Children = append(BuildTree())
		}
	}

	return Treenode

}
