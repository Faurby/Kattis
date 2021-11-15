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
	n := sc.Text()

	if strings.HasPrefix(n, "555") {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
