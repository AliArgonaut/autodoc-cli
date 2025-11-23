package utils

import (
	"cli/models"
)

func GetIgnoredFiles(cfg models.Config) []string {
	ignoredFiles := cfg.Ignore
	return ignoredFiles
}
