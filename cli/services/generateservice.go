package services

import (
	"fmt"
)

func GenerateService() {
	opts := GetIgnoredFilesService()
	tree := BuildTreeService()
	fmt.Println(tree)
}
