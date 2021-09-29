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

	t, _ := strconv.Atoi(sc.Text())

	for i := 1; i <= t; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Print(i % 100)
		}
		if i%3 == 0 {
			fmt.Print("øf")
		}
		if i%5 == 0 {
			fmt.Print("grynt")
		}

		fmt.Println()
	}
}
