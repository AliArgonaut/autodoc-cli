package services

import (
	"cli/services/utils"
	"fmt"
	"log"
	"strings"
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
	treeAsJSON := utils.ConvertTreeToJSON(node)

	allpaths := utils.TraverseTreeForPaths(node)
	langpaths := utils.GetLanguageFilesFromAllPaths(allpaths)
	text := utils.ReadLanguagePaths(langpaths)

	var sb strings.Builder
	sb.WriteString("program name: " + appName + "\n")
	sb.WriteString("program description: " + description + "\n")
	sb.WriteString("======FILE TREE=========" + "\n")
	sb.WriteString(string(treeAsJSON) + "\n")
	sb.WriteString("=======CODE (BY FILE)==========" + "\n")
	sb.WriteString(text)

	fmt.Println(sb.String())

}
