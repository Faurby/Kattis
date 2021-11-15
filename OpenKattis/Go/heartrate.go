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
		b, _ := strconv.ParseFloat(sc.Text(), 64)
		sc.Scan()
		p, _ := strconv.ParseFloat(sc.Text(), 64)

		bpm := (60 * b) / p
		min := bpm - (60 / p)
		max := bpm + (60 / p)
		fmt.Println(min, bpm, max)
	}
}
