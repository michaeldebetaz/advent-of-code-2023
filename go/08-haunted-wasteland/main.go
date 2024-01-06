package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	Key   string
	Left  string
	Right string
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

	var navigation []rune
	for _, r := range chunks[0] {
		navigation = append(navigation, r)
	}

	lines := strings.Split(chunks[1], "\n")

	var nodes []Node
	nodeMap := make(map[string]Node)

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		if len(parts) == 2 {
			key := parts[0]
			values := parts[1]
			values = strings.ReplaceAll(values, "(", "")
			values = strings.ReplaceAll(values, ")", "")
			data := strings.Split(values, ", ")
			left := data[0]
			right := data[1]
			node := Node{key, left, right}
			nodes = append(nodes, node)
			nodeMap[key] = node
		}
	}

	// fmt.Printf("Nodes: %v\n", nodes)

	// Part 1

	key := "AAA"
	steps := 0
	for i := 0; key != "ZZZ"; i = (i + 1) % len(navigation) {
		node := findNode(nodes, key)
		direction := navigation[i]
		key = getNextNodeKey(node, direction)
		// fmt.Printf("Node: %v, Direction: %v, Next Key: %v\n", node, string(direction), key)
		steps++
	}
	fmt.Printf("Steps (Part. 1): %v\n", steps)

	// Part 2

	var nodeList []Node
	for _, node := range nodes {
		if endsWith(node.Key, "A") {
			nodeList = append(nodeList, node)
		}
	}
	fmt.Printf("Node List: %v\n", nodeList)

	nodeSteps := make(map[string]int)
	for _, node := range nodeList {
		// fmt.Printf("Node: %v\n", node)
		found := false
		key := node.Key
		steps = 0
		for i := 0; found == false; i = (i + 1) % len(navigation) {
			n := nodeMap[key]
			direction := navigation[i]
			key = getNextNodeKey(n, direction)
			// fmt.Printf("Node: %v, Direction: %v, Next Key: %v\n", n, string(direction), key)
			steps++
			if endsWith(key, "Z") {
				found = true
			}
		}
		nodeSteps[node.Key] = steps
	}

	fmt.Printf("Node steps: %v\n", nodeSteps)

	maxSteps := nodeSteps[nodeList[0].Key]
	for _, steps := range nodeSteps {
		if steps > maxSteps {
			maxSteps = steps
		}
	}
	fmt.Printf("Min Steps: %v\n", maxSteps)

	result := maxSteps
	found := 0
	for i := 1; found < len(nodeSteps); i++ {
		v := maxSteps * i
		found = 0
		for _, steps := range nodeSteps {
			fmt.Printf("Steps: %v, V: %v\n", steps, v)
			if v%steps == 0 {
				found++
			}
			fmt.Printf("Found: %v\n", found)
		}
		result = v
	}
	fmt.Printf("Result: %v\n", result)
}

func findNode(nodes []Node, key string) Node {
	var node Node
	for _, n := range nodes {
		if n.Key == key {
			node = n
		}
	}
	return node
}

func getNextNodeKey(node Node, direction rune) string {
	var n string
	if direction == 'L' {
		n = node.Left
	} else {
		n = node.Right
	}
	return n
}

func endsWith(s string, char string) bool {
	return strings.HasSuffix(s, char)
}

func everyNodeKeyEndsWithZ(nodes []Node) bool {
	for _, node := range nodes {
		if !endsWith(node.Key, "Z") {
			return false
		}
	}
	return true
}
