package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		line := scanner.Text()
		text += line + "\n"
	}

	chunks := strings.Split(text, "\n\n")
	navigation := chunks[0]
	// network := chunks[1]

	lines := strings.Split(navigation, "\n")
	nodes := make(map[string]Node)

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		parent := parts[0]
		children := parts[1]
		children = strings.ReplaceAll(children, "(", "")
		children = strings.ReplaceAll(children, ")", "")
		data := strings.Split(children, ", ")
		left := data[0]
		right := data[1]
		nodes[parent] = Node{left, right}
	}

	fmt.Println(nodes)

}
