package utils

//import "log"

import "cli/models"
import "encoding/json"

func GoASTTOJSON(arr []models.ASTFile) []byte {
	//var bytes []byte
	jsonBytes, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	return jsonBytes
	//	for _, model := range arr {
	//		jsonBytes, err := json.Marshal(model)
	//
	//	/		if err != nil {
	//				log.Fatalln("failed to marshall json data")
	//			}
	//			bytes = append(bytes, jsonBytes...)
	//		}
	//		return bytes
}
