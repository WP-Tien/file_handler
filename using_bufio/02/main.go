package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	InputFileDir = "/file_handler/using_bufio/02/filetoread.txt"
)

func main() {

	var absPath = check_working_directory()

	// open input file for reading
	fi, err := os.Open(absPath + InputFileDir)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// Read/write
	fi2, err := os.OpenFile(absPath+InputFileDir, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	// Reading line by line
	// Use bufio.NewScanner to read the file line by line.
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println()

	// Reading with io.Reader
	// Use file.Read() to read data into a byte slice. It might require multiple calls to read the entire file.

	buffer := make([]byte, 1024)
	for {
		n, error := fi2.Read(buffer)
		if error == io.EOF {
			break
		}
		// if n == 0 {
		// 	break
		// }
		if error != nil {
			panic(error)
		}
		fmt.Println(string(buffer[:n]))
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
