package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := fmt.Sprintln(scanner.Text())
		fmt.Printf("Line: %s", line)

		line = strings.ReplaceAll(line, "two", "t2o")
		line = strings.ReplaceAll(line, "one", "o1e")
		line = strings.ReplaceAll(line, "three", "t3e")
		line = strings.ReplaceAll(line, "four", "4")
		line = strings.ReplaceAll(line, "five", "5e")
		line = strings.ReplaceAll(line, "six", "6")
		line = strings.ReplaceAll(line, "seven", "7n")
		line = strings.ReplaceAll(line, "nine", "n9e")
		line = strings.ReplaceAll(line, "eight", "e8t")

		fmt.Printf("Formatted line: %s\n", line)

		digits := make([]int, 0, len(line))

		for _, rune := range line {
			if isDigit(rune) {
				integer := int(rune - '0')
				digits = append(digits, integer)
			}
		}

		first := digits[0]
		last := digits[len(digits)-1]
		fmt.Printf("first: %v, last: %v\n", first, last)

		number := first*10 + last
		fmt.Printf("number: %v\n", number)

		sum = sum + number

		fmt.Println()
	}

	fmt.Printf("Sum: %v\n", sum)

}

func isDigit(rune rune) bool {
	return unicode.IsDigit(rune)
}
