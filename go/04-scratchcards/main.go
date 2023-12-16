package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	type Card struct {
		line            string
		winning_numbers []int64
		numbers         []int64
		value           int
	}

	cards := make(map[int]Card)

	for i := 0; scanner.Scan(); i++ {
		card_index := i + 1
		line := scanner.Text()

		substrings := strings.Split(line, ": ")
		// key := substrings[0]
		value := substrings[1]

		substr := strings.Split(value, " | ")
		re_numbers := regexp.MustCompile(`\d+`)
		left_numbers := re_numbers.FindAllString(substr[0], -1)
		right_numbers := re_numbers.FindAllString(substr[1], -1)

		winning_numbers := to_integers(left_numbers)
		numbers := to_integers(right_numbers)

		cards[card_index] = Card{line, winning_numbers, numbers, 0}
	}

	sum := 0
	for key, card := range cards {
		card.value = compute_card_value(card.winning_numbers, card.numbers)
		sum = sum + card.value
		fmt.Printf("Line %v: %s\n", key, card.line)
		fmt.Printf("Winning numbers: %v\n", card.winning_numbers)
		fmt.Printf("Numbers: %v\n", card.numbers)
		fmt.Printf("Value: %v\n", card.value)
		fmt.Printf("Sum: %v\n", sum)
	}
}

func to_integers(strings []string) []int64 {
	var numbers []int64
	for _, string := range strings {
		number, err := strconv.ParseInt(string, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	return numbers
}

func compute_card_value(winning_numbers []int64, numbers []int64) int {
	n_matches := 0

	for _, wn := range winning_numbers {
		for _, n := range numbers {
			if wn == n {
				n_matches++
			}
		}
	}

	if n_matches == 0 {
		return 0
	} else {
		return int(math.Pow(2, float64(n_matches-1)))
	}
}
