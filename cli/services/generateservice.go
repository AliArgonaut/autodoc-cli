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
	// description := utils.GetDescription(config)
	node := utils.BuildTree(cwd, name, ignoredFiles, "d")
	treeAsJSON := utils.ConvertTreeToJSON(node)

	goFiles, pyFiles := utils.DetectLanguages(node)

	//fmt.Println("========================================================================")
	//fmt.Println(description) //just to stop unused warning
	//fmt.Println("========================================================================")
	fmt.Println(treeAsJSON)
	fmt.Println("========================================================================")
	fmt.Println(goFiles)
	fmt.Println(pyFiles)
	//fmt.Println("==========================================================================")
}
