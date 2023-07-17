package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	deleteFiles()
}

func deleteFiles() {
	var directory string
	fmt.Scanln(&directory)
	for {
		readDirectory, _ := os.Open(directory)
		allFiles, _ := readDirectory.ReadDir(0)
		for f := range allFiles {
			file := allFiles[f]
			fileName := file.Name()
			filePath := directory + fileName
			err := os.RemoveAll(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		time.Sleep(30 * time.Minute)
	}
}
