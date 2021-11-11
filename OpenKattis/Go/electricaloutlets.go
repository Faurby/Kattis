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

	var output []int
	for i := 0; i < n; i++ {
		sum := 0
		sc.Scan()
		strips, _ := strconv.Atoi(sc.Text())
		for j := 0; j < strips; j++ {
			sc.Scan()
			outlets, _ := strconv.Atoi(sc.Text())
			if j == strips-1 {
				sum += outlets
			} else {
				sum += outlets - 1
			}
		}
		output = append(output, sum)
	}

	for _, value := range output {
		fmt.Println(value)
	}
}
