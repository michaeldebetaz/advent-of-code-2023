package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		lines = append(lines, line)

		re_gear := regexp.MustCompile(`\*`)
		gear_indexes := re_gear.FindAllStringIndex(line, -1)
		fmt.Printf("Gear indexes: %v\n", gear_indexes)
	}

	sum := int64(0)
	gears := make(map[string][]int64)
	for line_index, line := range lines {
		fmt.Printf("Line %v: %s\n", line_index, line)

		re_digits := regexp.MustCompile(`\d+`)

		numbers := re_digits.FindAllString(line, -1)
		number_indexes := re_digits.FindAllStringIndex(line, -1)

		for i, idxs := range number_indexes {
			number, err := strconv.ParseInt(numbers[i], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			is_part_number := int64(0)

			first := idxs[0]
			last := idxs[1] - 1
			has_prev_char := first > 0
			has_next_char := last < len(line)-1

			if has_prev_char {
				first--
			}

			if has_next_char {
				last++
			}

			// Previous line
			if line_index > 0 {
				prev_line_index := line_index - 1
				prev := lines[prev_line_index]
				chars := prev[first : last+1]
				is_part_number = has_symbol(chars, is_part_number)

				for i, rune := range chars {
					char := string(rune)
					fmt.Println(char)
					if char == "*" {
						gear_index := i + first
						append_number(gears, prev_line_index, gear_index, number)
					}
				}
			}

			// Current line
			if has_prev_char {
				prev_char := string(line[first])
				if prev_char != "." {
					is_part_number = 1
				}
				if prev_char == "*" {
					gear_index := first
					append_number(gears, line_index, gear_index, number)
				}
			}

			if has_next_char {
				next_char := string(line[last])
				if next_char != "." {
					is_part_number = 1
				}
				if next_char == "*" {
					gear_index := last
					append_number(gears, line_index, gear_index, number)
				}
			}

			// Next line
			if line_index < len(lines)-1 {
				next_line_index := line_index + 1
				next := lines[next_line_index]
				chars := next[first : last+1]
				is_part_number = has_symbol(chars, is_part_number)

				for i, rune := range chars {
					char := string(rune)
					fmt.Println(char)
					if char == "*" {
						gear_index := i + first
						append_number(gears, next_line_index, gear_index, number)
					}
				}
			}

			sum = sum + is_part_number*number
			fmt.Printf("number: %v; is part number = %v\n", number, is_part_number)
		}

		fmt.Printf("Number indexes: %v\n", number_indexes)
		fmt.Printf("Sum: %v\n", sum)
		fmt.Printf("Gears: %v\n", gears)

		sum_gears := int64(0)
		for _, numbers := range gears {
			if len(numbers) == 2 {
				sum_gears = sum_gears + (numbers[0] * numbers[1])
			}
		}

		fmt.Printf("Sum gears: %v\n", sum_gears)
	}
}

func has_symbol(s string, is_part_number int64) int64 {
	matched, err := regexp.Match(`^[.]+$`, []byte(s))
	if err != nil {
		log.Fatal(err)
	}
	if !matched {
		return 1
	}
	return is_part_number
}

func append_number(gears map[string][]int64, line_index int, gear_index int, number int64) {
	gear_key := fmt.Sprintf("%v - %v", line_index, gear_index)
	existing_numbers := gears[gear_key]
	gears[gear_key] = append(existing_numbers, number)
}
