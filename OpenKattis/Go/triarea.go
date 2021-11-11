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

	a, _ := strconv.ParseFloat(sc.Text(), 64)

	sc.Scan()
	b, _ := strconv.ParseFloat(sc.Text(), 64)

	fmt.Println((a * b) / 2)
}
