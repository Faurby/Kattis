package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()

	word := sc.Text()

	count := strings.Count(word, "e")

	var ees string
	for i := 0; i < 2*count; i++ {
		ees += "e"
	}
	fmt.Println("h" + ees + "y")
}
