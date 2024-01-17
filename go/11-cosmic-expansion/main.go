package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Printf("--- INIT ---\n\n")

	for _, line := range lines {
		fmt.Printf("%v\n", line)
	}

	fmt.Printf("\n--- PART. 1 ---\n\n")

}
