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

	for i := 0; i < n; i++ {
		sc.Scan()
		revenue, _ := strconv.Atoi(sc.Text())

		sc.Scan()
		expected, _ := strconv.Atoi(sc.Text())

		sc.Scan()
		cost, _ := strconv.Atoi(sc.Text())

		expcost := expected - cost

		if revenue > expcost {
			fmt.Println("do not advertise")
		} else if revenue == expcost {
			fmt.Println("does not matter")
		} else {
			fmt.Println("advertise")
		}
	}
}
