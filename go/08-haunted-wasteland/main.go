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

type Nodes map[string]Node

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

	var navigation []rune
	for _, r := range chunks[0] {
		navigation = append(navigation, r)
	}

	lines := strings.Split(chunks[1], "\n")
	nodes := make(Nodes)

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		if len(parts) == 2 {
			parent := parts[0]
			children := parts[1]
			children = strings.ReplaceAll(children, "(", "")
			children = strings.ReplaceAll(children, ")", "")
			data := strings.Split(children, ", ")
			left := data[0]
			right := data[1]
			nodes[parent] = Node{left, right}
		}
	}

	fmt.Println(nodes)

	steps := 0
	key := "AAA"
	for i := 0; key != "ZZZ"; i = (i + 1) % len(navigation) {
		direction := navigation[i]
		key = getNextKey(direction, nodes[key])
		fmt.Printf("Direction: %v, Key: %v\n", string(direction), key)
		steps++
	}
	fmt.Printf("Steps: %v\n", steps)
}

func getNextKey(r rune, node Node) string {
	var key string
	if r == 'L' {
		key = node.left
	} else {
		key = node.right
	}
	return key
}
