package services

import (
	"cli/services/utils"
	"fmt"
	"log"
)

func GenerateService() {

	cwd := utils.GetParentDirectory()
	name := utils.GetNameFromPath(cwd)

	config, err := utils.GetJSONData(cwd)
	if err != nil {
		log.Fatalln(err)
	}

	ignoredFiles := utils.GetIgnoredFiles(config)
	description := utils.GetDescription(config)
	appName := utils.GetAppName(config)
	node := utils.BuildTree(cwd, name, ignoredFiles, "d")

	allpaths := utils.TraverseTreeForPaths(node)
	langpaths := utils.GetLanguageFilesFromAllPaths(allpaths)
	text = utils.ReadLanguageFiles(langpaths)

	fmt.Println(description)
	fmt.Println(appName)
	fmt.Println(langpaths)
	// treeAsJSON := utils.ConvertTreeToJSON(node)

}
