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

	fmt.Println(resp.Body)

	defer resp.Body.Close()
	fmt.Printf("Response status: %d\n", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Response body length: %d\n", len(body))
	fmt.Printf("Response body preview: %s\n", string(body[:min(200, len(body))]))

	var response models.AgentResponseParams
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Failed to unmarshal: %v\nBody: %s\n", err, string(body))
	}

	fmt.Printf("Success field: %v\n", response.Success)
	fmt.Printf("Documentation length: %d\n", len(response.Documentation))

	if !response.Success {
		log.Fatalln("success not true")
	}

	// Print the documentation
	fmt.Println("\n=== DOCUMENTATION ===")
	fmt.Println(response.Documentation)
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
