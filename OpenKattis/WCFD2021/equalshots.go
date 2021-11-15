package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()

	drinkOne, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	drinkTwo, _ := strconv.Atoi(sc.Text())

	drinkOneSum := 0
	for i := 0; i < drinkOne; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())

		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		drinkOneSum += v * k
	}

	drinkTwoSum := 0
	for i := 0; i < drinkTwo; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())

		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		drinkTwoSum += v * k
	}

	if drinkOneSum == drinkTwoSum {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}
}
