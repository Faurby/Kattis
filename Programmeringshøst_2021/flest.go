package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	sc.Scan()

	//arrange words into map. String -> int
	animals := strings.Fields(sc.Text())
	var animalMap = make(map[string]int, len(animals))

	for _, i := range animals {
		animalMap[i] = 0
	}

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {

		sc.Scan()
		line := CleanLine(sc.Text())
		words := strings.Fields(line)

		mostOfAnimal := CountAnimals(words, animalMap)
		for key := range animalMap {
			count := 0
			for _, word := range words {
				if word == key {
					count++
				}
			}
			if count == mostOfAnimal {
				animalMap[key]++
			}
		}
	}

	for key, value := range animalMap {
		if value == greatestAnimal(animalMap) {
			fmt.Println(key)
		}
	}
}

func CleanLine(line string) string {
	result := strings.ToLower(line)
	result = strings.ReplaceAll(result, ",", "")
	result = strings.ReplaceAll(result, ".", " ")
	result = strings.ReplaceAll(result, "!", " ")
	// fmt.Println(result)
	return result
}

func CountAnimals(words []string, animalMap map[string]int) int {
	var greatest int
	for key := range animalMap {
		count := 0
		for _, word := range words {
			if word == key {
				count++
			}
		}
		if count > greatest {
			greatest = count
		}
	}
	return greatest
}

func greatestAnimal(animalMap map[string]int) int {
	greatest := 0
	for _, i := range animalMap {
		if i > greatest {
			greatest = i
		}
	}
	return greatest
}
