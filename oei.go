package oei

import (
	"log"
	"os/exec"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}

func AssertExec(command string, args ...string) string {
	output, err := exec.Command(command, args...).Output()
	outputString := string(output)
	log.Println(outputString)
	ErrorHandler(err)
	return outputString
}