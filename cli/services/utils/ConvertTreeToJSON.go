package utils

import (
	"cli/models"
	"encoding/json"
	"log"
)

func ConvertTreeToJSON(tree models.TreeNode) string {
	jsonBytes, err := json.Marshal(tree) //marshalIndent(tree, "", "   ") makes pretty JSON, but we want token efficiency
	if err != nil {
		log.Fatalln("failed to marshall json data")
	}
	return (string(jsonBytes))
}
