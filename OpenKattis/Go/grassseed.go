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
	cost, _ := strconv.ParseFloat(sc.Text(), 64)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sum := 0.0
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.ParseFloat(sc.Text(), 64)
		sc.Scan()
		b, _ := strconv.ParseFloat(sc.Text(), 64)

		sum += a * b
	}

	fmt.Println(sum * cost)
}
