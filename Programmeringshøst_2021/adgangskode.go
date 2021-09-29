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
	sc.Split(bufio.ScanWords)
	sc.Scan()

	// Arrange password length & animals.
	passLength, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	numberOfAnimals, _ := strconv.Atoi(sc.Text())
	var animals []string

	var result string

	for i := 0; i < numberOfAnimals; i++ {
		sc.Scan()

		animal := sc.Text()
		if len(animals) > 0 {
			for _, i := range animals {
				if len(i)+len(animal) == passLength {
					result = strings.ToLower(i + animal)
				}
			}
		}
		animals = append(animals, animal)
	}

	if len(result) != 0 {
		fmt.Println(result)
	} else {
		fmt.Println("*umuligt*")
	}
}
