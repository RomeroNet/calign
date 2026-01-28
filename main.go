package main

import (
	"calign/internal/file"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: calign <file>")
		os.Exit(1)
	}

	path := os.Args[1]

	if path[0] != '/' {
		workingDir, err := os.Getwd()

		if err != nil {
			log.Fatal(err)
		}

		if path[0] == '.' && path[1] == '/' {
			path = path[2:]
		}

		path = workingDir + "/" + path
	}

	f, err := file.FromPath(path)

	if err != nil {
		log.Fatal(err)
	}

	err = f.Process()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("caligned " + path)
}
