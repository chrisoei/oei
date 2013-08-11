package oei

import (
	"log"
	"os/exec"
	"path/filepath"
)

// Simple error handler that prints the error and exits the program
func ErrorHandler(err error) {
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}

// Run a command with arguments, print the output, exit if any errors, return result
func AssertExec(command string, args ...string) string {
	output, err := exec.Command(command, args...).Output()
	outputString := string(output)
	log.Println(outputString)
	ErrorHandler(err)
	return outputString
}

// Return the full filename without the file extension
func FilenameWithoutExt(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
