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
	// from here youd need to figure out what other context to send and send it over to the agents as an octet stream
}
