package services

import (
	"cli/models"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func BuildTreeService() string {

	cwd := GetCWD()
	name := GetName(cwd)
	tree := BuildTree(cwd, name)
	treeAsString := convertToJson(tree)
	return treeAsString
}

func BuildTree(cwd string, name string) models.TreeNode {
	Treenode := models.TreeNode{}
	Treenode.Name = name
	Treenode.Path = cwd

	entries, err := os.ReadDir(cwd)
	if err != nil {
		log.Fatalln(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			Treenode.Children = append(Treenode.Children, BuildTree(filepath.Join(cwd, entry.Name()), entry.Name()))
		} else {
			Treenode.Children = append(Treenode.Children, models.TreeNode{
				Name: entry.Name(),
				Path: filepath.Join(cwd, entry.Name()),
			})
		}
	}

	return Treenode

}

func GetCWD() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("couldnt get parent directory")
	}
	return string(wd)
}

func GetName(d string) string {
	return string(filepath.Base(d))
}

func convertToJson(tree models.TreeNode) string {
	jsonBytes, err := json.MarshalIndent(tree, "", "    ")
	if err != nil {
		log.Fatalln("failed to marshall json data")
	}
	return (string(jsonBytes))
}
