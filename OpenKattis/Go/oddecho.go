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

	var result []string
	for i := 0; i < n; i++ {
		sc.Scan()
		if i%2 == 0 {
			result = append(result, sc.Text())
		}
	}

	for _, value := range result {
		fmt.Println(value)
	}
}
