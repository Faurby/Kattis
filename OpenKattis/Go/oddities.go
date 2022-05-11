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
		number, _ := strconv.Atoi(sc.Text())
		if number%2 == 0 {
			fmt.Printf("%d is even\n", number)
		} else {
			fmt.Printf("%d is odd\n", number)
		}
	}
}
