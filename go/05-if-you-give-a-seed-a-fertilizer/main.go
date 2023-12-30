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

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var text string

	for scanner.Scan() {
		text = text + scanner.Text() + "\n"
	}

	chunks := strings.Split(text, "\n\n")

	seeds := getSeeds(chunks[0])
	fmt.Printf("Seeds: %v\n\n", seeds)
	maps := getMaps(chunks, seeds)

	var lowestLocation int
	for i, seed := range seeds {
		location := getLocation(seed, maps)
		if i == 0 || location < lowestLocation {
			lowestLocation = location
		}
	}
	fmt.Printf("Lowest location (pt. 1): %v\n", lowestLocation)

	// seedPairs := getSeedPairs(chunks[0])

}

type SeedPair struct {
	Start int
	Range int
}

func getSeeds(chunk string) []int {
	data := getSeedData(chunk)
	var seeds []int
	for _, value := range data {
		seed := parseInt(value)
		seeds = append(seeds, seed)
	}

	return seeds

}

func getSeedPairs(chunk string) []SeedPair {
	data := getSeedData(chunk)
	var seedPairs []SeedPair
	for i := 0; i < len(data); i = i + 2 {
		start := parseInt(data[i])
		r := parseInt(data[i+1])
		seed := SeedPair{start, r}
		seedPairs = append(seedPairs, seed)
	}

	return seedPairs
}

func getSeedData(chunk string) []string {
	values := strings.Replace(chunk, "seeds: ", "", 1)
	data := strings.Split(values, " ")
	return data
}

func getMaps(chunks []string, seeds []int) Maps {
	seedToSoilMap := getMap(chunks[1])
	fmt.Printf("Seed to soil map: %v\n\n", seedToSoilMap)
	soilToFertilizerMap := getMap(chunks[2])
	fmt.Printf("Soil to fetilizer map: %v\n\n", soilToFertilizerMap)
	fertilizerToWaterMap := getMap(chunks[3])
	fmt.Printf("Fertilizer to water map: %v\n\n", fertilizerToWaterMap)
	waterToLightMap := getMap(chunks[4])
	fmt.Printf("Water to light map: %v\n\n", waterToLightMap)
	lightToTemperatureMap := getMap(chunks[5])
	fmt.Printf("Light to temparature map: %v\n\n", lightToTemperatureMap)
	temperatureToHumidityMap := getMap(chunks[6])
	fmt.Printf("Temperature to humidity map: %v\n\n", temperatureToHumidityMap)
	humidityToLocationMap := getMap(chunks[7])
	fmt.Printf("Humidity to location map: %v\n\n", humidityToLocationMap)

	maps := Maps{seedToSoilMap, soilToFertilizerMap, fertilizerToWaterMap, waterToLightMap, lightToTemperatureMap, temperatureToHumidityMap, humidityToLocationMap}

	return maps
}

func getLocation(seed int, maps Maps) int {
	soil := getDestination(seed, maps.SeedTosoilMap)
	fertilizer := getDestination(soil, maps.SoilToFertilizerMap)
	water := getDestination(fertilizer, maps.FertilizerToWaterMap)
	light := getDestination(water, maps.WaterToLightMap)
	temperature := getDestination(light, maps.LightToTemperatureMap)
	humidity := getDestination(temperature, maps.TemperatureToHumidityMap)
	location := getDestination(humidity, maps.HumidityToLocationMap)
	return location
}

func getMap(chunk string) []Coordinates {
	lines := strings.Split(chunk, "\n")
	data := lines[1:]

	var coordinates []Coordinates
	for _, line := range data {
		if line != "" {
			values := strings.Split(line, " ")
			source := parseInt(values[0])
			destination := parseInt(values[1])
			r := parseInt(values[2])
			coordinate := Coordinates{source, destination, r}
			coordinates = append(coordinates, coordinate)
		}
	}

	return coordinates
}

func getDestination(source int, coordinates []Coordinates) int {
	destination := source
	for _, c := range coordinates {
		MIN := c.Source
		MAX := MIN + c.Range
		if MIN <= source && source <= MAX {
			delta := source - MIN
			destination = c.Destination + delta
		}
	}
	return destination
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
