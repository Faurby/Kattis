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

	if strings.Contains(word, "ss") {
		fmt.Println("hiss")
	} else {
		fmt.Println("no hiss")
	}
}
