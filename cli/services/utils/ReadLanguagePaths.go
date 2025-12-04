package utils

import (
	"cli/models"
	"os"
	"regexp"
	"strings"
)

func ReadLanguagePaths(paths []string) ([]models.CodeFile, error) {

	var result = []models.CodeFile{}

	spaceRegex := regexp.MustCompile(`\s+`)
	for _, v := range paths {
		fileContent, err := os.ReadFile(v)
		if err != nil {
			return nil, err
		}

		code := strings.TrimSpace(spaceRegex.ReplaceAllString(string(fileContent), " "))
		result = append(result, models.CodeFile{
			Filename: v,
			Code:     code,
		})

	}
	return result, nil
}
