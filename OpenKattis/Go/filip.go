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
	a, _ := strconv.Atoi(Reverse(sc.Text()))
	sc.Scan()
	b, _ := strconv.Atoi(Reverse(sc.Text()))

	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
