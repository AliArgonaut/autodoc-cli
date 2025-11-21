package services

import "fmt"
import "cli/models"

func AutodocInitService() {
	fmt.Println("auto doc init service init")
	models.NewConfigTemplate()
}
