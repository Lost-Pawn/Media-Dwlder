package main

import (
	"fmt"
	"bytes"
	"os/exec"
)

func main() {
	var out bytes.Buffer
	var stderr bytes.Buffer
	URL := "Your url link" // define a func for toolName
	cmd := exec.Command(toolName, URL) // "-g", "-f", "-o", "-d" 
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("Error Output: " + stderr.String())
		fmt.Println("Execution Error: " + err.Error())
	} else {
		fmt.Println("Successfully Uploaded.." + out.String())
	}
}
