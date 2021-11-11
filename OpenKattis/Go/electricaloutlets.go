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

	var sum int
	var output []int
	for i := 0; i < n; i++ {
		sc.Scan()
		strips, _ := strconv.Atoi(sc.Text())
		fmt.Printf("Strips: %v", strips)
		for j := 0; j < strips; j++ {
			sc.Scan()
			outlets, _ := strconv.Atoi(sc.Text())
			fmt.Printf("outlets: %v", outlets)
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
