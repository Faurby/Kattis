package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	N, _ := strconv.ParseFloat(sc.Text(), 64)

	sc.Scan()
	K, _ := strconv.ParseFloat(sc.Text(), 64)

	sc.Scan()
	H, _ := strconv.ParseFloat(sc.Text(), 64)

	if r := N - (K * H); r < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(math.Ceil(r / H))
	}
}
