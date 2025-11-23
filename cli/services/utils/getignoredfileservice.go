package utils

import "os"

func GetIgnoredFilesArrayService(d string) []string {
	entries, err := os.ReadDir(d)
	//return errors.New()

}
