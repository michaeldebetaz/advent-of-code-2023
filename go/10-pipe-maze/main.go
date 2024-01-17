package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Pipe struct {
	Row  int
	Col  int
	Rune rune
}

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

	start := Pipe{}
	for row, line := range lines {
		for col, rune := range line {
			if rune == 'S' {
				start = Pipe{Row: row, Col: col, Rune: rune}
				break
			}
		}
	}
	fmt.Printf("Start: %v\n", start)

	pipes := []Pipe{}
	pipe := start
	nextPipe := Pipe{}
	prevPipe := Pipe{}
	found := false
	steps := 0
	for found == false {
		pipes = append(pipes, pipe)
		nextPipe, prevPipe, found = walk(pipe, prevPipe, lines)
		// fmt.Printf("Next: %v, Prev: %v, Found: %v\n", nextPipe, prevPipe, found)
		pipe = nextPipe
		steps++

		if pipe.Row == 0 && pipe.Col == 0 {
			break
		}
	}

	furthest := float32(steps / 2)

	fmt.Printf("Start Pipe: %v, End Pipe: %v, Steps: %v, Furthest: %v\n", start, pipe, steps, furthest)

	fmt.Print("\n--- PART. 2 ---\n\n")

	sum := 0

	tiles := [][2]int{}
	for i, line := range lines {
		for j := range line {
			curr := [2]int{i, j}
			intersections := 0
			if !isPipe(curr, pipes) {
				for k := 0; k != j; k++ {
					point := [2]int{i, k}
					rune := rune(lines[i][k])
					if isPipe(point, pipes) && (rune == '|' || rune == 'J' || rune == 'L') {
						intersections++
					}
				}
				if intersections%2 == 1 {
					tiles = append(tiles, curr)
					sum++
				}
			}
		}
	}

	for i, line := range lines {
		for j := range line {
			point := [2]int{i, j}
			if isPipe(point, pipes) {
				fmt.Printf("🟦")
			} else if isTile(point, tiles) {
				fmt.Printf("🟩")
			} else {
				fmt.Printf("🟥")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Sum: %v\n", sum)

}

func isPipe(point [2]int, pipes []Pipe) bool {
	for _, p := range pipes {
		if p.Row == point[0] && p.Col == point[1] {
			return true
		}
	}
	return false
}

func isTile(point [2]int, tiles [][2]int) bool {
	for _, t := range tiles {
		if t[0] == point[0] && t[1] == point[1] {
			return true
		}
	}
	return false
}

func walk(pipe Pipe, prevPipe Pipe, lines []string) (Pipe, Pipe, bool) {
	nextPipes := []Pipe{}
	nextPipe := Pipe{}
	found := false

	topRow := pipe.Row - 1
	bottomRow := pipe.Row + 1
	leftCol := pipe.Col - 1
	rightCol := pipe.Col + 1

	// Top
	if pipe.Row > 0 && (pipe.Rune == '|' || pipe.Rune == 'L' || pipe.Rune == 'J' || pipe.Rune == 'S') {
		top := rune(lines[topRow][pipe.Col])
		if top == '|' || top == '7' || top == 'F' || top == 'S' {
			nextPipe := Pipe{Row: topRow, Col: pipe.Col, Rune: top}
			nextPipes = append(nextPipes, nextPipe)
		}
	}

	// Right
	if pipe.Col < len(lines[0])-1 && (pipe.Rune == '-' || pipe.Rune == 'L' || pipe.Rune == 'F' || pipe.Rune == 'S') {
		right := rune(lines[pipe.Row][rightCol])
		if right == '-' || right == '7' || right == 'J' || right == 'S' {
			nextPipe := Pipe{Row: pipe.Row, Col: rightCol, Rune: right}
			nextPipes = append(nextPipes, nextPipe)
		}
	}

	// Bottom
	if pipe.Row < len(lines)-1 && (pipe.Rune == '|' || pipe.Rune == '7' || pipe.Rune == 'F' || pipe.Rune == 'S') {
		bottom := rune(lines[bottomRow][pipe.Col])
		if bottom == '|' || bottom == 'J' || bottom == 'L' || bottom == 'S' {
			nextPipe := Pipe{Row: bottomRow, Col: pipe.Col, Rune: bottom}
			nextPipes = append(nextPipes, nextPipe)
		}
	}

	// Left
	if pipe.Col > 0 && (pipe.Rune == '-' || pipe.Rune == '7' || pipe.Rune == 'J' || pipe.Rune == 'S') {
		left := rune(lines[pipe.Row][leftCol])
		if left == '-' || left == 'F' || left == 'L' || left == 'S' {
			nextPipe := Pipe{Row: pipe.Row, Col: leftCol, Rune: left}
			nextPipes = append(nextPipes, nextPipe)
		}
	}

	for _, p := range nextPipes {
		isPrev := p.Row == prevPipe.Row && p.Col == prevPipe.Col && p.Rune == prevPipe.Rune
		if !isPrev {
			nextPipe = p
		}
	}

	found = nextPipe.Rune == 'S'

	// fmt.Printf("Pipe: %v, Next Pipes: %v, Selected: %v, Prev Pipe: %v\n", pipe, nextPipes, nextPipe, prevPipe)

	return nextPipe, pipe, found
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
