package services

import (
	"fmt"
	"io"
	"net/http"
)

func HealthCheck() error {
	fmt.Println("checking health of auto-doc backend servers...")
	resp, err := http.Get("http://localhost:8000/api/health")
	if err != nil {
		fmt.Println("unable to query auto-doc backend servers...try again later")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("unable to read auto-doc backend server response...try again later")
	}
	fmt.Println(string(body))
	return nil
}
