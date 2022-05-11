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

	number, _ := strconv.Atoi(sc.Text())

	for {
		if isHassadNumber(number) {
			fmt.Println(number)
			break
		} else {
			number += 1
		}
	}
}

func isHassadNumber(number int) bool {
	remainder := 0
	sumResult := 0
	testnumber := number
	for testnumber != 0 {
		remainder = testnumber % 10
		sumResult += remainder
		testnumber = testnumber / 10
	}

	return number%sumResult == 0
}
