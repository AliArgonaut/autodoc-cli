package utils

import (
	"log"
	"os"
)

func GetParentDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("couldnt get parent directory")
	}
	return string(wd)
}
