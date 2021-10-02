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

	n, _ := strconv.Atoi(sc.Text())

	var sum, valid, weight int
	weightError := -2
	for i := 0; i < n; i++ {
		sc.Scan()
		weight, _ = strconv.Atoi(sc.Text())

		if i > weightError+1 {
			if weight >= 10 && weight <= 2000 {
				sum += weight
				valid++
			} else {
				weightError = i
			}
		}
	}
	fmt.Printf("%f", float64(sum)/float64(valid))
}
