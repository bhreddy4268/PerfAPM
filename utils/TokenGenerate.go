package utils

import (
	"fmt"
	"os/exec"
)

func main() {
	// Command to run the shell script
	cmd := exec.Command("/bin/bash", "Generate_Token.sh") // Replace "your_script.sh" with your script's path

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output of the shell script if needed
	fmt.Println(string(output))
}

