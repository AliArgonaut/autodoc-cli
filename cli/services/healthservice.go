package services

import (
	"fmt"
	"io"
	"net/http"
)

func HealthCheck() error {
	fmt.Println("checking health of autodoc backend server...")
	resp, err := http.Get("http://localhost:8000/api/health")
	if err != nil {
		fmt.Println("ERROR: could not reach autodoc backend server. did you run 'autodoc start'?")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR: unable to read auto-doc backend server response...re-clone the autodoc git repo and rebuild if you can't identify the problem")
	}
	fmt.Println(string(body))
	return nil
}
