package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReadLanguagePaths(paths []string) string {
	var sb strings.Builder
	spaceRegex := regexp.MustCompile(`\s+`)
	for _, v := range paths {
		fileContent, err := os.ReadFile(v)
		if err != nil {
			fmt.Println("error reading file")
		}
		sb.WriteString(fmt.Sprintf("========== %s ==========", v))

		codeLine := strings.TrimSpace(spaceRegex.ReplaceAllString(string(fileContent), " "))

		sb.WriteString("\n")
		sb.Write([]byte(codeLine))
		sb.WriteString("\n")
	}
	return sb.String()
}
