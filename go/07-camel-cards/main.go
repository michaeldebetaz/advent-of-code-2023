package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	CardsString string
	Cards       []int
	CardCount   map[int]int
	Bid         int
	Type        int
	Rank        int
}

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// getResult(file, 1)
	getResult(file, 2)
}

func getResult(file *os.File, part int) {
	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		substrings := strings.Split(line, " ")

		var cards []int
		cardsString := substrings[0]
		for _, rune := range cardsString {
			card := getStrenght(rune, part)
			cards = append(cards, card)
		}

		bid := parseInt(substrings[1])

		hands = append(hands, Hand{CardsString: cardsString, Cards: cards, Bid: bid})
	}

	for i := range hands {
		hand := &hands[i]
		hand.CardCount = getHandCardCount(hand.Cards)
		fmt.Printf("Cards: %v, Count: %v\n", hand.Cards, hand.CardCount)
		hand.Type = getHandType(*hand)
	}

	hands = sortHands(hands, part)

	sum := 0
	for i := range hands {
		hand := &hands[i]
		rank := i + 1
		hand.Rank = rank
		sum += hand.Bid * rank
	}

	for _, hand := range hands {
		fmt.Printf("Cards: %v, Type: %v, Rank: %v\n", hand.Cards, hand.Type, hand.Rank)
	}

	fmt.Printf("Sum (part %v): %v\n", part, sum)
}

func getStrenght(r rune, part int) int {
	if r == 'A' {
		return 14
	}
	if r == 'K' {
		return 13
	}
	if r == 'Q' {
		return 12
	}
	if r == 'J' {
		if part == 1 {
			return 11
		}
		if part == 2 {
			return 1
		}
	}
	if r == 'T' {
		return 10
	}
	return int(r - '0')
}

func getHandCardCount(cards []int) map[int]int {
	m := make(map[int]int)
	maxCount := 0
	var maxCountCard int
	for _, card := range cards {
		if _, ok := m[card]; ok {
			m[card]++
		} else {
			m[card] = 1
		}
		if m[card] > maxCount && card != 1 {
			maxCount = m[card]
			maxCountCard = card
		}
	}

	if _, ok := m[1]; ok {
		count := m[1]
		m[maxCountCard] = m[maxCountCard] + count
		delete(m, 1)
	}

	return m
}

func getHandType(hand Hand) int {
	var counts []int
	for _, count := range hand.CardCount {
		counts = append(counts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	t := 1
	if counts[0] == 5 {
		t = 7
	}
	if counts[0] == 4 {
		t = 6
	}
	if counts[0] == 3 {
		if counts[1] == 2 {
			t = 5
		} else {
			t = 4
		}
	}
	if counts[0] == 2 {
		if counts[1] == 2 {
			t = 3
		} else {
			t = 2
		}
	}

	return t
}

func sortHands(hands []Hand, part int) []Hand {
	sort.SliceStable(hands, func(index, nextIndex int) bool {
		prevHand := hands[index]
		nextHand := hands[nextIndex]
		return prevHand.Type < nextHand.Type
	})

	sort.SliceStable(hands, func(index, nextIndex int) bool {
		prevHand := hands[index]
		nextHand := hands[nextIndex]
		if prevHand.Type == nextHand.Type {
			prevHandCards := replaceCardLetters(prevHand.CardsString, part)
			nextHandCards := replaceCardLetters(nextHand.CardsString, part)
			return prevHandCards < nextHandCards
		}
		return false
	})
	return hands
}

func replaceCardLetters(s string, part int) string {
	s = strings.ReplaceAll(s, "A", "Z")
	s = strings.ReplaceAll(s, "K", "Y")
	s = strings.ReplaceAll(s, "Q", "X")
	if part == 1 {
		s = strings.ReplaceAll(s, "J", "W")
	}
	if part == 2 {
		s = strings.ReplaceAll(s, "J", "1")
	}
	s = strings.ReplaceAll(s, "T", "V")
	return s
}

func switchHands(hand *Hand, nextHand *Hand) {
	tmp := *hand
	*hand = *nextHand
	*nextHand = tmp
}

func parseInt[V string | rune](v V) int {
	n, err := parseIntFromStringOrRune(v)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func parseIntFromStringOrRune[V string | rune](v V) (int, error) {
	var i interface{} = v
	if s, ok := i.(string); ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return n, nil
	}
	if r, ok := i.(rune); ok {
		n := int(r - '0')
		return n, nil
	}
	error := errors.New("Cannot parse type " + reflect.TypeOf(v).String())
	return 0, error
}
