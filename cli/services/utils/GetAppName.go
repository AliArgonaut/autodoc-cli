package utils

import (
	"cli/models"
)

func GetAppName(cfg models.Config) string {
	name := cfg.Name
	return name

}
