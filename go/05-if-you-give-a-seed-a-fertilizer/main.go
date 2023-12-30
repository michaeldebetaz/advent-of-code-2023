package main

import (
	"bufio"
	"log"
	"os"
<<<<<<< Updated upstream
	"strconv"
	"strings"
=======
>>>>>>> Stashed changes
)

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var text string

	for scanner.Scan() {
		text = text + scanner.Text() + "\n"
	}

<<<<<<< Updated upstream
	chunks := strings.Split(text, "\n\n")

	seeds := getSeeds(chunks[0])

	var seedToSoilMap []Map

}

type Seed struct {
	Start int
	Range int
}

func getSeeds(chunk string) []Seed {
	values := strings.Replace(chunk, "seeds: ", "", 1)
	data := strings.Split(values, " ")

	var seeds []Seed
	for i := 0; i < len(data); i = i + 2 {
		start := parseInt(data[i])
		r := parseInt(data[i+1])
		seed := Seed{start, r}
		seeds = append(seeds, seed)
	}

	return seeds
}

func getMaps(chunks []string) Maps {
	maps := make(Maps)

}

func getMap(chunk string) []Coordinates {
	lines := strings.Split(chunk, "\n")
	data := lines[1:]

	var coordinates []Coordinates
	for _, line := range data {
		values := strings.Split(line, " ")
		source := parseInt(values[0])
		destination := parseInt(values[1])
		r := parseInt(values[2])
		coordinate := Coordinates{source, destination, r}
		coordinates = append(coordinates, coordinate)
	}

	return coordinates
}

type Coordinates struct {
	Destination int
	Source      int
	Range       int
}

type Maps struct {
	SeedTosoilMap            []Coordinates
	SoilToFertilizerMap      []Coordinates
	FertilizerToWaterMap     []Coordinates
	WaterToLightMap          []Coordinates
	LightToTemperatureMap    []Coordinates
	TemperatureToHumidityMap []Coordinates
	HumidityToLocationMap    []Coordinates
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
=======
	var seedToSoilMap []Map

}
>>>>>>> Stashed changes
