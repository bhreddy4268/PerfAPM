package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Asset struct {
	InfaToken string `json:"infaToken"`
	// Add other fields if present in the JSON response
}

func main() {
	// Create an HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Adjust the timeout duration as needed
	}

	// Make the HTTP request
	resp, err := client.Get("http://irl75cmgperf06:7085/ldmadmin/web.isp/loginJs")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Create a JSON decoder for the response body
	decoder := json.NewDecoder(resp.Body)

	// Initialize an empty Asset struct
	var obj Asset

	// Decode the JSON response incrementally
	for decoder.More() {
		if err := decoder.Decode(&obj); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		// Process obj or extract the infaToken here if needed
		infaToken := obj.InfaToken
		fmt.Println("InfaToken:", infaToken)
	}
}
