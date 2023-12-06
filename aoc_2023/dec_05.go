package main

import (
	"math"
	"strings"
)

type Mapping struct {
	destStart   int
	sourceStart int
	length      int
}

func (m *Mapping) inSourceRange(x int) bool {
	start := m.sourceStart
	end := m.sourceStart + m.length - 1
	return x >= start && x <= end
}

func (m *Mapping) inDestRange(x int) bool {
	return x >= m.destStart && x <= m.destStart+m.length
}

func (m *Mapping) getDest() []int {
	sourceRange := []int{m.destStart, m.destStart + m.length - 1}

	return sourceRange
}

func (m *Mapping) getSource() []int {
	sourceRange := []int{m.sourceStart, m.sourceStart + m.length - 1}

	return sourceRange
}

func solvePart02() {
	//var input = ReadFileIntoArray("resources/dec_05/example.txt")
	var input = ReadFileIntoArray("resources/dec_05/input.txt")

	seeds := strings.Split(input[0], ":")
	seedNumbers := StringToNumberOfIntegers(seeds[1])

	var start []int
	var length []int

	for idx := 0; idx < len(seedNumbers); idx += 2 {
		start = append(start, seedNumbers[idx])
		length = append(length, seedNumbers[idx+1])
	}

	i := 1
	for input[i] != "seed-to-soil map:" {
		i++
	}

	i++

	var seedToSoilMappings []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		seedToSoilMappings = append(seedToSoilMappings, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "soil-to-fertilizer map:" {
		i++
	}
	i++

	var soilToFertilizerMap []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		soilToFertilizerMap = append(soilToFertilizerMap, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "fertilizer-to-water map:" {
		i++
	}
	i++

	var fertilizerToWater []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		fertilizerToWater = append(fertilizerToWater, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "water-to-light map:" {
		i++
	}
	i++

	var waterToLightMap []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		waterToLightMap = append(waterToLightMap, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "light-to-temperature map:" {
		i++
	}
	i++

	var lightToTemperature []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		lightToTemperature = append(lightToTemperature, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "temperature-to-humidity map:" {
		i++
	}
	i++

	var temperatureToHumidity []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		temperatureToHumidity = append(temperatureToHumidity, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "humidity-to-location map:" {
		i++
	}
	i++

	var humidityToLocation []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		humidityToLocation = append(humidityToLocation, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	min := math.MaxInt
	numberOfSeeds := 0
	for k := 0; k < len(start); k++ {
		startSeeds := start[k]
		endSeeds := start[k] + length[k]
		for l := startSeeds; l < endSeeds; l++ {
			soilNumber := l
			numberOfSeeds++
			for _, soil := range seedToSoilMappings {
				if soil.inSourceRange(l) {
					soilNumber = findDest(soil, l)
				}
			}
			//println("Soil: " + strconv.Itoa(soilNumber))

			fertilizer := soilNumber
			for _, soil := range soilToFertilizerMap {
				if soil.inSourceRange(soilNumber) {
					fertilizer = findDest(soil, soilNumber)
				}
			}
			//println("Fert: " + strconv.Itoa(fertilizer))

			water := fertilizer
			for _, soil := range fertilizerToWater {
				if soil.inSourceRange(fertilizer) {
					water = findDest(soil, fertilizer)
				}
			}
			/*println("Water: " + strconv.Itoa(water))*/

			light := water
			for _, soil := range waterToLightMap {
				if soil.inSourceRange(water) {
					light = findDest(soil, water)
				}
			}

			/*println("Light: " + strconv.Itoa(light))*/

			temperature := light
			for _, soil := range lightToTemperature {
				if soil.inSourceRange(light) {
					temperature = findDest(soil, light)
				}
			}
			/*println("Temp: " + strconv.Itoa(temperature))*/

			humidity := temperature
			for _, soil := range temperatureToHumidity {
				if soil.inSourceRange(temperature) {
					humidity = findDest(soil, temperature)
				}
			}

			/*println("Hum: " + strconv.Itoa(humidity))*/

			location := humidity
			for _, soil := range humidityToLocation {
				if soil.inSourceRange(humidity) {
					location = findDest(soil, humidity)
				}
			}

			//println("Location: " + strconv.Itoa(location))

			if location < min {
				min = location
			}
		}
		println(numberOfSeeds)
	}

	println(min)
}

func part01() {
	var input = ReadFileIntoArray("resources/dec_05/example.txt")
	//var input = ReadFileIntoArray("resources/dec_05/input.txt")

	seeds := strings.Split(input[0], ":")
	seedNumbers := StringToNumberOfIntegers(seeds[1])

	var start []int
	var length []int

	for idx := 0; idx < len(seedNumbers); idx += 2 {
		start = append(start, seedNumbers[idx])
		length = append(length, seedNumbers[idx+1])
	}

	i := 1
	for input[i] != "seed-to-soil map:" {
		i++
	}

	i++

	var seedToSoilMappings []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		seedToSoilMappings = append(seedToSoilMappings, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "soil-to-fertilizer map:" {
		i++
	}
	i++

	var soilToFertilizerMap []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		soilToFertilizerMap = append(soilToFertilizerMap, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "fertilizer-to-water map:" {
		i++
	}
	i++

	var fertilizerToWater []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		fertilizerToWater = append(fertilizerToWater, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "water-to-light map:" {
		i++
	}
	i++

	var waterToLightMap []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		waterToLightMap = append(waterToLightMap, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "light-to-temperature map:" {
		i++
	}
	i++

	var lightToTemperature []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		lightToTemperature = append(lightToTemperature, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "temperature-to-humidity map:" {
		i++
	}
	i++

	var temperatureToHumidity []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		temperatureToHumidity = append(temperatureToHumidity, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	for input[i] != "humidity-to-location map:" {
		i++
	}
	i++

	var humidityToLocation []Mapping
	for i < len(input) && input[i] != "" {
		ranges := StringToNumberOfIntegers(input[i])
		humidityToLocation = append(humidityToLocation, Mapping{sourceStart: ranges[1], destStart: ranges[0], length: ranges[2]})
		i++
	}

	min := math.MaxInt

	for k := 0; k < len(start); k++ {
		startSeeds := start[k]
		endSeeds := start[k] + length[k]
		for l := startSeeds; k < endSeeds; l++ {
		}
	}

	for _, seed := range seedNumbers {
		soilNumber := seed
		for _, soil := range seedToSoilMappings {
			if soil.inSourceRange(seed) {
				soilNumber = findDest(soil, seed)
			}
		}
		//println("Soil: " + strconv.Itoa(soilNumber))

		fertilizer := soilNumber
		for _, soil := range soilToFertilizerMap {
			if soil.inSourceRange(soilNumber) {
				fertilizer = findDest(soil, soilNumber)
			}
		}
		//println("Fert: " + strconv.Itoa(fertilizer))

		water := fertilizer
		for _, soil := range fertilizerToWater {
			if soil.inSourceRange(fertilizer) {
				water = findDest(soil, fertilizer)
			}
		}
		/*println("Water: " + strconv.Itoa(water))*/

		light := water
		for _, soil := range waterToLightMap {
			if soil.inSourceRange(water) {
				light = findDest(soil, water)
			}
		}

		/*println("Light: " + strconv.Itoa(light))*/

		temperature := light
		for _, soil := range lightToTemperature {
			if soil.inSourceRange(light) {
				temperature = findDest(soil, light)
			}
		}
		/*println("Temp: " + strconv.Itoa(temperature))*/

		humidity := temperature
		for _, soil := range temperatureToHumidity {
			if soil.inSourceRange(temperature) {
				humidity = findDest(soil, temperature)
			}
		}

		/*println("Hum: " + strconv.Itoa(humidity))*/

		location := humidity
		for _, soil := range humidityToLocation {
			if soil.inSourceRange(humidity) {
				location = findDest(soil, humidity)
			}
		}

		//println("Location: " + strconv.Itoa(location))

		if location < min {
			min = location
		}
	}

	println(min)
}

func findDest(soil Mapping, seed int) int {
	start := soil.sourceStart
	//end := soil.sourceStart + soil.length - 1

	index := seed - start

	return soil.destStart + index
}
