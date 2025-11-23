package utils

import (
	"cli/models"
	"encoding/json"
	"errors"
	"log"
	"os"
)

func GetJSONData(d string) (models.Config, error) {
	var cfgStruct = models.Config{}
	entries, err := os.ReadDir(d)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range entries {
		if file.Name() == "autodoc_config.json" {
			jsonData, err := os.ReadFile("autodoc_config.json")
			if err != nil {
				log.Fatalln(err)
			}

			json.Unmarshal(jsonData, &cfgStruct)
			return cfgStruct, nil
		}
	}
	return cfgStruct, errors.New("couldnt find autodoc config")
}
