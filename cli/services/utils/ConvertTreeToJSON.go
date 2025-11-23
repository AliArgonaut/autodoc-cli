package utils

import (
	"cli/models"
	"encoding/json"
	"log"
)

func ConvertTreeToJSON(tree models.TreeNode) string {
	jsonBytes, err := json.MarshalIndent(tree, "", "    ")
	if err != nil {
		log.Fatalln("failed to marshall json data")
	}
	return (string(jsonBytes))
}
