package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var multipliers = map[string]int{
	"hundrede":  100,
	"tusinde":   1_000,
	"million":   1_000_000,
	"millioner": 1_000_000,
}

var numericValue = map[string]int{
	"nul":        0,
	"en":         1,
	"et":         1,
	"to":         2,
	"tre":        3,
	"fire":       4,
	"fem":        5,
	"seks":       6,
	"syv":        7,
	"otte":       8,
	"ni":         9,
	"ti":         10,
	"elleve":     11,
	"tolv":       12,
	"tretten":    13,
	"fjorten":    14,
	"femten":     15,
	"seksten":    16,
	"sytten":     17,
	"atten":      18,
	"nitten":     19,
	"tyve":       20,
	"tredive":    30,
	"fyrre":      40,
	"halvtreds":  50,
	"tres":       60,
	"halvfjerds": 70,
	"firs":       80,
	"halvfems":   90,
}

var sum, number, multiplier int

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	sc.Scan()

	words := strings.Fields(sc.Text())

	var tempSum, previousTempSum int
	if len(words) == 1 {
		fmt.Println(parseNumericalValue(words[0]))
	} else {
		for _, word := range words {

			if strings.Contains(word, "og") {
				fmt.Println("A")
				number = parseNumericalValue(word)
				multiplier = 1
				tempSum = number * multiplier
			} else {
				fmt.Println("B")
				if _number := parseNumericalValue(word); _number > 0 {
					number = _number
					fmt.Println(number)
				} else if _multiplier := parseMultiplier(word); _multiplier > 0 {
					fmt.Println(multiplier)
					multiplier = _multiplier
					tempSum = number * multiplier
				} else if _multiplier == 0 {
					multiplier = 1
				}
			}
			if !(tempSum == previousTempSum) {
				sum += tempSum
				previousTempSum = tempSum
			}
			fmt.Printf("number %v and multiplier %v and sum %v \n", number, multiplier, sum)
		}
		fmt.Println(sum)
	}
}

func parseNumericalValue(word string) int {
	if strings.Contains(word, "og") {
		numbers := strings.Split(word, "og")
		return numericValue[numbers[0]] + numericValue[numbers[1]]
	} else {
		return numericValue[word]
	}
}

func parseMultiplier(word string) int {
	return multipliers[word]
}
