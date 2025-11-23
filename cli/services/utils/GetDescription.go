package utils

import (
	"cli/models"
)

func GetDescription(cfg models.Config) string {
	description := cfg.Description
	return description
}
