package services

import (
	"cli/services/utils"
	"fmt"
)

func GenerateService() {

	cwd := utils.GetParentDirectory()
	// fmt.Println("cwd::::::::::  " + cwd)
	name := utils.GetNameFromPath(cwd)
	// fmt.Println("name::::::::::  " + name)

	ignoredFiles, err := utils.GetIgnoredFiles(cwd)
	if err != nil {
		fmt.Println(err)
	}
	//	fmt.Println("ignoredFiles::::::::::  " + ignoredFiles[1])
	node := utils.BuildTree(cwd, name, ignoredFiles)
	json := utils.ConvertTreeToJSON(node)
	// fmt.Println(json)
}
