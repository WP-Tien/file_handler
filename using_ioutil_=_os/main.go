package main

import (
	"log"
	"os"
)

const (
	InputFileDir  = "/file_handler/using_ioutil_=_os/input.txt"
	OutputFileDir = "/file_handler/using_ioutil_=_os/output.txt"
)

func main() {
	var absPath string
	absPath = check_working_directory()

	// read the whole file at once
	b, err := os.ReadFile(absPath + InputFileDir)
	if err != nil {
		panic(err)
	}

	// write the whole body at once
	err = os.WriteFile(absPath+OutputFileDir, b, 0644)
	if err != nil {
		panic(err)
	}
}

/*
Check Working Directory: Ensure the program is run from the directory containing the file or use the os.Getwd() function to print the current working directory.
*/
func check_working_directory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}
