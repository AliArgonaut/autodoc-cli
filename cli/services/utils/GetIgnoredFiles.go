package utils

import (
	"cli/models"
	"encoding/json"
	"errors"
	"log"
	"os"
)

func GetIgnoredFiles(d string) ([]string, error) {

	var cfgStruct = models.Config{}

	entries, err := os.ReadDir(d)
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range entries {
		if entry.Name() == "autodoc_config.json" {
			jsonData, err := os.ReadFile("autodoc_config.json")
			if err != nil {
				log.Fatalln(err)
			}

			json.Unmarshal(jsonData, &cfgStruct)
			ignoredFiles := cfgStruct.Ignore
			return ignoredFiles, nil
		}
	}
	return nil, errors.New("couldnt find autodoc_config")
}
