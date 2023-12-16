package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum_part_1 := 0
	sum_part_2 := 0

	for scanner.Scan() {
		line := fmt.Sprintf(scanner.Text())
		fmt.Printf("Line: %s\n", line)

		s := strings.Split(line, ": ")

		id_str := strings.TrimPrefix(s[0], "Game ")
		id, err := strconv.ParseInt(id_str, 10, 8)
		if err != nil {
			log.Fatal(err)
		}

		possible := int64(1)
		power := int64(0)

		games := strings.Split(s[1], "; ")

		var max = make(map[string]int64)
		max["red"] = 0
		max["green"] = 0
		max["blue"] = 0

		for _, game := range games {
			subsets := strings.Split(game, ", ")

			for _, subset := range subsets {
				res := strings.Split(subset, " ")
				n, err := strconv.ParseInt(res[0], 10, 8)
				if err != nil {
					log.Fatal(err)
				}

				switch color := res[1]; color {
				case "red":
					if n > MAX_RED {
						possible = 0
					}
					if n > max["red"] {
						max["red"] = n
					}
				case "green":
					if n > MAX_GREEN {
						possible = 0
					}
					if n > max["green"] {
						max["green"] = n
					}
				case "blue":
					if n > MAX_BLUE {
						possible = 0
					}
					if n > max["blue"] {
						max["blue"] = n
					}
				}
			}
		}

		fmt.Printf("Possible: %v\n", possible)
		sum_part_1 = sum_part_1 + int(possible*id)
		fmt.Printf("Sum (Part 1): %v\n", sum_part_1)

		power = max["red"] * max["green"] * max["blue"]
		fmt.Printf("Power (Part 2): %v\n", power)
		sum_part_2 = sum_part_2 + int(power)
		fmt.Printf("Sum (Part 2): %v\n", sum_part_2)
	}
	fmt.Printf("Result (Part 1): %v\n", sum_part_1)
	fmt.Printf("Result (Part_2): %v\n", sum_part_2)
}
