package oei

import (
	"log"
	"os/exec"
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