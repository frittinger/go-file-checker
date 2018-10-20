package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func readFileNames(filename string) []string {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var filenames []string

	for scanner.Scan() {
		filenames = append(filenames, scanner.Text())
	}

	return filenames
}

func main() {
	var files []string

	root := "."
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	var searchFilenames = readFileNames("data.txt")
	fmt.Println(searchFilenames)

	var existingfilenames []string
	for _, file := range files {
		var reducedFilename = filepath.Base(file)
		existingfilenames = append(existingfilenames, reducedFilename)
	}
	fmt.Println("--------------")
	fmt.Println(existingfilenames)

	// search for files that do not exist in the existing files

	var missingfilenames []string
	for i := 0; i < len(searchFilenames); i++ {
		exists := false
		for j := 0; j < len(existingfilenames); j++ {
			if searchFilenames[i] == existingfilenames[j] {
				exists = true
				break
			}
		}
		if !exists {
			missingfilenames = append(missingfilenames, searchFilenames[i])
		}
	}

	fmt.Println("######################")
	fmt.Println(missingfilenames)

}
