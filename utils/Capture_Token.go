package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Read the content from the text file
	content, err := ioutil.ReadFile("response.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the file content to a string
	responseString := string(content)

	// Find and extract the InfaToken value using string manipulation
	startIndex := strings.Index(responseString, "infaToken: \"")
	if startIndex != -1 {
		startIndex += len("infaToken: \"")
		endIndex := strings.Index(responseString[startIndex:], "\"")
		if endIndex != -1 {
			infaToken := responseString[startIndex : startIndex+endIndex]
			// Print only the dynamic InfaToken value without the label
			fmt.Println(infaToken)
		}
	} else {
		fmt.Println("InfaToken not found")
	}
}
