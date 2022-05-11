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
	low, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	high, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	resultSum, _ := strconv.Atoi(sc.Text())

	for i := low; i <= high; i++ {
		if sumDigits(i) == resultSum {
			fmt.Println(i)
			break
		}
	}
	for i := high; i >= low; i-- {
		if sumDigits(i) == resultSum {
			fmt.Println(i)
			break
		}
	}
}

func sumDigits(number int) int {
	sumResult := 0
	for number != 0 {
		remainder := number % 10
		sumResult += remainder
		number = number / 10
	}

	return sumResult
}
