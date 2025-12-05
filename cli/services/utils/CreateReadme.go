package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateReadme(cwd string, text string) {
	filename := "README.md"
	filepath := filepath.Join(cwd, filename)
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("error creating file")
		return
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("couldn't write to file")
		return
	}
	fmt.Println("File created at:", filepath)
}
