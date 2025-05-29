package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

const (
	InputFileDir  = "/file_handler/using_bufio/01/input.txt"
	OutputFIleDir = "/file_handler/using_bufio/01/output.txt"
)

func main() {
	var absPath string
	absPath = check_working_directory()

	// open input file
	fi, err := os.Open(absPath + InputFileDir)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(fi)

	// open output file
	fo, err := os.Create(absPath + OutputFIleDir)
	if err != nil {
		panic(err)
	}
	// close fo on ecit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err = w.Flush(); err != nil {
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
