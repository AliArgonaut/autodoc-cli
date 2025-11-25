package services

import (
	"cli/models"
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

	goFiles, _ := utils.DetectLanguages(node)
	print(goFiles)
	var goAst []models.ASTFile = utils.GetGoASTFiles(cwd, goFiles)
	astAsJson := utils.GoASTTOJSON(goAst)
	fmt.Println(treeAsJSON)
	fmt.Println("========================================================================")
	fmt.Println(astAsJson)
}
