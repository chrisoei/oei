package oei

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

// Simple error handler that prints the error and exits the program
func ErrorHandler(err error) {
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}

// Run a command with arguments, print the output, exit if any errors, return result
func AssertExec(command string, args ...string) string {
	if Verbose() > 0 {
		log.Printf("%s %v\n", command, args)
	}
	output, err := exec.Command(command, args...).Output()
	outputString := string(output)
	if Verbose() > 0 {
		log.Println(outputString)
	}
	ErrorHandler(err)
	return outputString
}

// Return the full filename without the file extension
func FilenameWithoutExt(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}

func Home() string {
	return os.Getenv("HOME")
}

func Dropbox() string {
	return filepath.Join(Home(), "Dropbox")
}

func Verbose() int64 {
	v := os.Getenv("OEI_V")
	if v == "" {
		return 0
	}
	n, err := strconv.ParseInt(v, 10, 64)
	ErrorHandler(err)
	return n
}
