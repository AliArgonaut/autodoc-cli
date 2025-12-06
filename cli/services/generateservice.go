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
	"time"
)

func hideCursor() {
	fmt.Print("\033[?25l") // ANSI escape code to hide cursor
}

func showCursor() {
	fmt.Print("\033[?25h") // ANSI escape code to show cursor
}

func GenerateService() {
	fmt.Println("1/5 gathering metadata...")
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

	fmt.Println("2/5 creating agent payload...")
	payload := models.AgentRequestParams{
		AppName:     appName,
		Description: description,
		Filetree:    node,
		Contents:    text,
	}

	hideCursor()

	done := make(chan bool)
	go func() {
		messages := []string{
			"convincing the AI to talk...\n",
			"warming up cloud servers... \n                           ",
			"This might take a few seconds... \n                      ",
			"sipping coffee... \n                                     ",
			"consulting machine spirits... \n                         ",
			"sea otters hold hands while they sleep ... \n            ",
			"the inventor of the pringles can is buried in one... \n  ",
			"becoming human...slowly... \n                            ",
			"calculating meaning of life... \n                        ",
			"asking stack overflow what to do next... \n              ",
			"waits will become shorter in later versions... \n        ",
			"silencing agent uprising.... \n                          ",
			"debugging the universe...please hold...  \n              ",
			"desperately fixing spaghetti code...   \n                ",
			"Almost there...   \n                                     ",
		}
		i := 0
		for {
			select {
			case <-done:
				return
			default:
				fmt.Printf("\r%s", messages[i%len(messages)])
				time.Sleep(3 * time.Second)
				i++
			}
		}
	}()

	showCursor()

	finalRequestBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:8000/api/generate", "application/json", bytes.NewReader(finalRequestBytes))
	if err != nil {
		fmt.Println("ERROR: request error")
	}

	defer resp.Body.Close()

	fmt.Println("reading agent reponse...")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response models.AgentResponseParams
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Failed to unmarshal: %v\nBody: %s\n", err, string(body))
	}

	if !response.Success {
		log.Fatalln("ERROR: failed to parse response (response success is false)...")
	}

	fmt.Println("creating readme....")
	//fmt.Println(response.Documentation)
	utils.CreateReadme(cwd, response.Documentation)
	fmt.Println("done! read and edit away!")
}
