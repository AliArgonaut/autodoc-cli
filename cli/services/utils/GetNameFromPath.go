package utils

import "path/filepath"

func GetNameFromPath(d string) string {
	return string(filepath.Base(d))
}
