package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
	Ways     []int
}

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var lines [][]string
	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(line, -1)
		lines = append(lines, matches)
	}

	var data [][]int
	for _, line := range lines {
		var numbers []int
		for _, s := range line {
			n := parseInt(s)
			numbers = append(numbers, n)
		}
		data = append(data, numbers)
	}
	times := data[0]
	distances := data[1]

	var races []Race
	for i, time := range times {
		distance := distances[i]
		var ways []int
		race := Race{time, distance, ways}
		races = append(races, race)
	}

	for i := range races {
		race := &races[i]
		time := race.Time
		for hold := 1; hold <= time; hold++ {
			distance := getDistance(hold, time)
			if distance > race.Distance {
				race.Ways = append(race.Ways, hold)
			}
		}
	}
	fmt.Printf("Races: %v\n\n", races)

	product := 1
	for _, race := range races {
		product = product * len(race.Ways)
	}
	fmt.Printf("Product: %v\n\n", product)

	t := join(lines[0])
	d := join(lines[1])
	var ways []int
	race := Race{parseInt(t), parseInt(d), ways}
	fmt.Printf("Race: %v\n\n", race)

	time := race.Time
	for hold := 1; hold <= time; hold++ {
		distance := getDistance(hold, race.Time)
		if distance > race.Distance {
			first := hold
			last := time - first
			ways := last - first + 1
			fmt.Printf("First: %v | Last: %v | Ways: %v\n\n", first, last, ways)
			break
		}
	}

}

func getDistance(hold int, time int) int {
	speed := hold
	remainingTime := time - hold
	return speed * remainingTime
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func join(array []string) string {
	return strings.Join(array, "")
}
