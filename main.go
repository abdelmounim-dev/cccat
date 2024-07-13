package main

import (
	"fmt"
	"log"
	"os"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("not implemented yet")
		return
	}
	for i := 1; i < len(os.Args); i++ {
		fileName := os.Args[i]
		if !fileExists(fileName) {
			log.Fatalf("file %s doesn't exist", fileName)
		}
		content, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal("error reading file: ", err)
		}
		fmt.Printf("%s\n", string(content))
	}
}
