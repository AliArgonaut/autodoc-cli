package utils

import "log"

import "cli/models"
import "encoding/json"

func GoASTTOJSON(arr []models.ASTFile) string {
	var results = ""

	for _, model := range arr {
		jsonBytes, err := json.Marshal(model)
		if err != nil {
			log.Fatalln("failed to marshall json data")
		}
		results += string(jsonBytes)
	}
	return results
}
