package utils

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
	"os/exec"
)

// Declare the infaToken as a global variable
var infaToken []byte // Change the type to []byte

func TokenGenerate() {
    // Command to run the shell script
    cmd := exec.Command("/bin/bash", "Generate_Token.sh") // Replace "your_script.sh" with your script's path

    // Execute the command
    var err error
    infaToken, err = cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the output of the shell script if needed
    fmt.Println(string(infaToken))
}

func LDMLogin(URL string, client *http.Client, User string, Pass string, infaToken string) (*http.Response, *http.Client) {
    // Convert the global infaToken variable to a string here
    infaTokenStr := string(infaToken)

    values := map[string]string{"user": User, "password": Pass, "namespace": "Native", "infaToken": infaTokenStr}

    jsonValue, _ := json.Marshal(values)
    // pass the values to the request's body

    resp, err := client.Post(URL, "application/json", bytes.NewBuffer(jsonValue))

    if err != nil {
        fmt.Println(resp)
    }

    return resp, client
}
