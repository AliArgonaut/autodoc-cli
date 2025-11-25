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
	treeAsJSON := utils.ConvertTreeToJSON(node)

	goFiles, _ := utils.DetectLanguages(node)
	//	print(goFiles)
	var goAst []models.ASTFile = utils.GetGoASTFiles(cwd, goFiles)
	astAsJson := utils.GoASTTOJSON(goAst)
	//fmt.Println(string(astAsJson))

	var agentRequestParams = models.AgentRequestParams{
		AppName:     appName,
		Description: description,
		Filetree:    json.RawMessage(treeAsJSON),
		AST:         json.RawMessage(astAsJson),
	}

	req, err := json.Marshal(agentRequestParams)
	if err != nil {
		panic(err)
	}

	print(req)
	sendAgentParamsToAgentLayer(req)

}

func sendAgentParamsToAgentLayer(req []byte) error {
	//fmt.Println("making post request with gathered context...")
	resp, err := http.Post("http://localhost:8000/api/generate", "application/json", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println("unable to query auto-doc backend server...try again later (maybe our agents are sleeping)")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("unable to read auto-doc backend server response...try again later")
	}
	fmt.Println(string(body))
	return nil

}
