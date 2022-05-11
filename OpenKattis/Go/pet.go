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

	var max, counter, sum int

	for i := 1; i <= 5; i++ {
		sum = 0
		for j := 0; j < 4; j++ {
			sc.Scan()
			number, _ := strconv.Atoi(sc.Text())
			sum += number
		}

		if sum > max {
			max = sum
			counter = i
		}
	}
	fmt.Printf("%d %d", counter, max)
}
