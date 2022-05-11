package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dominantPoint = map[string]int{
	"A": 11,
	"K": 4,
	"Q": 3,
	"J": 20,
	"T": 10,
	"9": 14,
	"8": 0,
	"7": 0,
}

var nondominantPoint = map[string]int{
	"A": 11,
	"K": 4,
	"Q": 3,
	"J": 2,
	"T": 10,
	"9": 0,
	"8": 0,
	"7": 0,
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	hands, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	dominantSuit := sc.Text()

	sum := 0
	for i := 0; i < 4*hands; i++ {
		sc.Scan()
		cards := strings.Split(sc.Text(), "")

		if cards[1] == dominantSuit {
			sum += dominantPoint[cards[0]]
		} else {
			sum += nondominantPoint[cards[0]]
		}
	}

	fmt.Print(sum)
}
