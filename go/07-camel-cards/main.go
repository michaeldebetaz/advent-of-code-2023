package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

	}
}
