package services

import (
	"cli/services/utils"
	"fmt"
	"log"
)

func GenerateService() {

	cwd := utils.GetParentDirectory()  //obvious
	name := utils.GetNameFromPath(cwd) //

	config, err := utils.GetJSONData(cwd)
	if err != nil {
		log.Fatalln(err)
	}

	ignoredFiles := utils.GetIgnoredFiles(config)
	description := utils.GetDescription(config)
	node := utils.BuildTree(cwd, name, ignoredFiles)
	json := utils.ConvertTreeToJSON(node)
	fmt.Println(json)
	// from here youd need to figure out what other context to send and send it over to the agents as an octet stream
}
