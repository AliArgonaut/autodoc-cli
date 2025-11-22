package services

import (
	"fmt"
)

func GenerateService() {
	tree := BuildTreeService()
	fmt.Println(tree)
}
