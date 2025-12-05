package services

import (
	"bytes"
	"cli/models"
	"cli/services/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	//	treeAsJSON := utils.ConvertTreeToJSON(node)

	allpaths := utils.TraverseTreeForPaths(node)
	langpaths := utils.GetLanguageFilesFromAllPaths(allpaths)
	text, err := utils.ReadLanguagePaths(langpaths)
	if err != nil {
		log.Fatalln(err)
	}

	payload := models.AgentRequestParams{
		AppName:     appName,
		Description: description,
		Filetree:    node,
		Contents:    text,
	}

	finalRequestBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:8000/api/generate", "application/json", bytes.NewReader(finalRequestBytes))
	if err != nil {
		fmt.Println("request error")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response models.AgentResponseParams
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Failed to unmarshal: %v\nBody: %s\n", err, string(body))
	}

	if !response.Success {
		log.Fatalln("success not true")
	}

	//fmt.Println(response.Documentation)
	utils.CreateReadme(cwd, response.Documentation)
}
