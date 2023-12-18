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
		copies          int
	}

	var cards []Card

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		substrings := strings.Split(line, ": ")
		substr := strings.Split(substrings[1], " | ")
		re_numbers := regexp.MustCompile(`\d+`)
		left_numbers := re_numbers.FindAllString(substr[0], -1)
		right_numbers := re_numbers.FindAllString(substr[1], -1)

		winning_numbers := to_integers(left_numbers)
		numbers := to_integers(right_numbers)

		card := Card{line, winning_numbers, numbers, 0, 1}
		cards = append(cards, card)
	}

	sum := 0
	for index, card := range cards {
		matches, value := get_matches_and_card_value(card.winning_numbers, card.numbers)

		card.value = value
		sum = sum + card.value
		cards[index] = card

		for i := 0; i < matches; i++ {
			card_index := index + i + 1
			if card_index < len(cards) {
				c := &cards[card_index]
				c.copies = c.copies + (1 * card.copies)
			}
		}

		fmt.Println(card.line)
		fmt.Printf("Winning numbers: %v\n", card.winning_numbers)
		fmt.Printf("Numbers: %v\n", card.numbers)
		fmt.Printf("Value: %v\n", card.value)
		fmt.Printf("Sum: %v\n", sum)
	}

	copies := 0
	for index, card := range cards {
		copies = copies + card.copies
		fmt.Printf("Card %v: value = %v; copies = %v\n", index+1, card.value, card.copies)
	}
	fmt.Printf("Total copies = %v\n", copies)
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

func get_matches_and_card_value(winning_numbers []int64, numbers []int64) (int, int) {
	matches := 0

	for _, wn := range winning_numbers {
		for _, n := range numbers {
			if wn == n {
				matches++
			}
		}
	}

	if matches == 0 {
		return matches, 0
	} else {
		return matches, int(math.Pow(2, float64(matches-1)))
	}
}
