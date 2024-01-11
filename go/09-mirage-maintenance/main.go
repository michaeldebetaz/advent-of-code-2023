package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var histories [][][]int
	for scanner.Scan() {
		line := scanner.Text()
		substrings := strings.Split(line, " ")

		var sequences [][]int
		var sequence []int
		for _, substring := range substrings {
			n := parseInt(substring)
			sequence = append(sequence, n)
		}
		sequences = append(sequences, sequence)

		found := false
		for found == false {
			sequence, found = getNextSequence(sequence)
			sequences = append(sequences, sequence)
		}

		histories = append(histories, sequences)
	}
	fmt.Printf("--- INIT ---\n\n")
	fmt.Printf("First history: %v\n", histories[0])

	findNextValues(histories)

	findPrevValues(histories)
}

func findNextValues(histories [][][]int) {
	var nextValues []int
	sum := 0

	for _, history := range histories {
		lastIndex := len(history) - 1
		for i := lastIndex; i > 0; i-- {
			currSeq := history[i]
			currSeqLastIndex := len(currSeq) - 1
			currSeqLast := currSeq[currSeqLastIndex]

			prevSeq := history[i-1]
			prevSeqLastIndex := len(prevSeq) - 1
			prevSeqLast := prevSeq[prevSeqLastIndex]
			prevSeqNext := prevSeqLast + currSeqLast
			prevSeq = append(prevSeq, prevSeqNext)
			history[i-1] = prevSeq
		}
		firstSeq := history[0]
		lastIndex = len(firstSeq) - 1
		nextValue := firstSeq[lastIndex]
		nextValues = append(nextValues, nextValue)
		sum += nextValue
	}

	fmt.Printf("\n--- PART. 1 ---\n\n")
	fmt.Printf("First history: %v\n", histories[0])
	fmt.Printf("Next values: %v\n", nextValues)
	fmt.Printf("Sum: %v\n", sum)
}

func findPrevValues(histories [][][]int) {
	var prevValues []int
	sum := 0

	for _, history := range histories {
		firstIndex := len(history) - 1
		for i := firstIndex; i > 0; i-- {
			currSeq := history[i]
			currSeqFirst := currSeq[0]

			prevSeq := history[i-1]
			prevSeqFirst := prevSeq[0]
			prevSeqPrev := prevSeqFirst - currSeqFirst
			prevSeq = append([]int{prevSeqPrev}, prevSeq...)
			history[i-1] = prevSeq
		}
		firstSeq := history[0]
		prevValue := firstSeq[0]
		prevValues = append(prevValues, prevValue)
		sum += prevValue
	}

	fmt.Printf("\n--- PART. 2 ---\n\n")
	fmt.Printf("First history: %v\n", histories[0])
	fmt.Printf("Prev values: %v\n", prevValues)
	fmt.Printf("Sum: %v\n", sum)
}

func getNextSequence(sequence []int) ([]int, bool) {
	var tmp []int
	for i := 0; i < len(sequence)-1; i++ {
		curr := sequence[i]
		next := sequence[i+1]
		diff := next - curr
		tmp = append(tmp, diff)
	}

	found := true
	for _, n := range tmp {
		if n != 0 {
			found = false
		}
	}

	return tmp, found
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
